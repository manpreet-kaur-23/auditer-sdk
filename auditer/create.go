package auditer

import (
	"github.com/manpreet-kaur-23/auditer-sdk/kafka"
	"gorm.io/gorm"
)

func beforeCreateCallback(db *gorm.DB) {
	_ = kafka.SendKafkaMessage("audit_log_test", "Before create callback called")
}

func afterCreateCallback(db *gorm.DB) {
	_ = kafka.SendKafkaMessage("audit_log_test", "After create callback called")
}
