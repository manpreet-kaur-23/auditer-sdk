package auditer

import (
	"github.com/manpreet-kaur-23/auditer-sdk/kafka"
	"gorm.io/gorm"
)

func beforeDeleteCallback(db *gorm.DB) {
	_ = kafka.SendKafkaMessage("audit_log_test", "Before delete callback called")
}

func afterDeleteCallback(db *gorm.DB) {
	_ = kafka.SendKafkaMessage("audit_log_test", "After delete callback called")
}
