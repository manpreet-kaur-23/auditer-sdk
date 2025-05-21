package auditer

import (
	"log"
	"reflect"

	"github.com/manpreet-kaur-23/auditer-sdk/kafka"
	"gorm.io/gorm"
)

func afterCreateCallback(db *gorm.DB) {
	// Convert the created record to a map[string]interface{}
	log.Println("=======CREATE-AUDITER======")
	record := map[string]interface{}{}
	if err := db.Statement.Parse(db.Statement.Dest); err == nil {
		// Use GORM's built-in function to convert struct to map
		for _, field := range db.Statement.Schema.Fields {
			value, _ := field.ValueOf(db.Statement.Context, reflect.ValueOf(db.Statement.Dest))
			record[field.Name] = value
		}
	}

	GlobalAuditLog.Record = record
	GlobalAuditLog.Event = "create"
	GlobalAuditLog.Resource = db.Statement.Schema.Name

	_ = kafka.SendAuditLogToKafka("audit_log_test", GlobalAuditLog)
}
