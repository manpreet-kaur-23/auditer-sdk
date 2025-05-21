package auditer

import (
	"log"

	"github.com/manpreet-kaur-23/auditer-sdk/kafka"
	"gorm.io/gorm"
)

func beforeDeleteCallback(db *gorm.DB) {
	log.Println("========DELETE-AUDITER========")
	// Get the model type
	model := db.Statement.Model

	// Prepare a slice to hold the records about to be deleted
	records := make(map[string]interface{}, 0)

	// Query the records that match the current delete filter
	// This uses the current statement's conditions
	tx := db.Session(&gorm.Session{}).Model(model)
	if db.Statement.Unscoped {
		tx = tx.Unscoped()
	}
	if err := tx.Where(db.Statement.Clauses["WHERE"].Expression).Find(&records).Error; err != nil {
		log.Printf("Failed to fetch records for audit before delete: %v", err)
	} else {
		GlobalAuditLog.Record = records
		GlobalAuditLog.Event = "Delete"
		GlobalAuditLog.Resource = db.Statement.Table
		_ = kafka.SendAuditLogToKafka("audit_log_test", GlobalAuditLog)
	}
}
