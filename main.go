package main

import (
	"github.com/manpreet-kaur-23/auditer-sdk/auditer"
	"gorm.io/gorm"
)

func InitSDK(db *gorm.DB) {

	db.Use(&auditer.AuditerPlugin{})
}
