package auditer

import (
	"gorm.io/gorm"
	"github.com/manpreet-kaur-23/auditer-sdk/kafka"
)

func beforeUpdateCallback(db *gorm.DB) {
	_ = kafka.SendKafkaMessage("auditer-callbacks", "Before update callback called")
}

func afterUpdateCallback(db *gorm.DB) {
	_ = kafka.SendKafkaMessage("auditer-callbacks", "After update callback called")
}
