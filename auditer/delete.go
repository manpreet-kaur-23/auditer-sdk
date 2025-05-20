package auditer

import (
	"gorm.io/gorm"
	"github.com/manpreet-kaur-23/auditer-sdk/kafka"
)

func beforeDeleteCallback(db *gorm.DB) {
	_ = kafka.SendKafkaMessage("auditer-callbacks", "Before delete callback called")
}

func afterDeleteCallback(db *gorm.DB) {
	_ = kafka.SendKafkaMessage("auditer-callbacks", "After delete callback called")
}
