package auditer

import (
	"github.com/manpreet-kaur-23/auditer-sdk/kafka"
	"gorm.io/gorm"
)

func beforeCreateCallback(db *gorm.DB) {
	_ = kafka.SendAuditLogToKafka("audit_log_test", GlobalAuditLog)
}

func afterCreateCallback(db *gorm.DB) {
	_ = kafka.SendAuditLogToKafka("audit_log_test", GlobalAuditLog)
}
