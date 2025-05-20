package auditer

import (
	"gorm.io/gorm"
	"github.com/manpreet-kaur-23/auditer-sdk/kafka"
)

func beforeCreateCallback(db *gorm.DB) {
	_ = kafka.SendKafkaMessage("auditer-callbacks", "Before create callback called")
}

func afterCreateCallback(db *gorm.DB) {
	_ = kafka.SendKafkaMessage("auditer-callbacks", "After create callback called")
}
