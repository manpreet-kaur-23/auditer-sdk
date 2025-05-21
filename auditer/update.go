package auditer

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/manpreet-kaur-23/auditer-sdk/kafka"
	"gorm.io/gorm"
)

const instanceKeyOld = "model_before_update"
const instanceKeyNew = "model_after_update"

func beforeUpdateCallback(db *gorm.DB) {
	// _ = kafka.SendAuditLogToKafka("audit_log_test", GlobalAuditLog)
	if db.Statement.Schema == nil {
		return
	}
	// Store the original value in the GORM instance context
	db.InstanceSet(instanceKeyOld, db.Statement.ReflectValue.Interface())
}

func afterUpdateCallback(db *gorm.DB) {
	log.Println("========UPDATE-AUDITER========")
	oldVal, okOld := db.InstanceGet(instanceKeyOld)
	newVal := db.Statement.ReflectValue.Interface()

	oldMap, ok1 := structToMap(oldVal)
	newMap, ok2 := structToMap(newVal)
	if !okOld || !ok1 || !ok2 {
		return
	}

	recordChanges := make(map[string][]interface{})
	for k, newV := range newMap {
		oldV := oldMap[k]
		if _, exists := oldMap[k]; !exists || !reflect.DeepEqual(oldV, newV) {
			recordChanges[k] = []interface{}{oldV, newV}
		}
	}

	GlobalAuditLog.Record = newMap
	GlobalAuditLog.RecordChanges = recordChanges
	GlobalAuditLog.Event = "update"
	GlobalAuditLog.Resource = db.Statement.Schema.Name

	_ = kafka.SendAuditLogToKafka("audit_log_test", GlobalAuditLog)
}

func structToMap(obj interface{}) (map[string]interface{}, bool) {
	var m map[string]interface{}
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, false
	}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, false
	}
	return m, true
}
