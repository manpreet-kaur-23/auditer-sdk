package auditer

import (
	"github.com/manpreet-kaur-23/auditer-sdk/kafka"
	"gorm.io/gorm"
)

func beforeUpdateCallback(db *gorm.DB) {
	_ = kafka.SendKafkaMessage("audit_log_test", "Before update callback called")
}

func afterUpdateCallback(db *gorm.DB) {
	_ = kafka.SendKafkaMessage("audit_log_test", "After update callback called")
}
