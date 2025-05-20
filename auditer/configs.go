package auditer

import (
	"log"

	"gorm.io/gorm"
)

type AuditerPlugin struct{}

func (p *AuditerPlugin) Name() string {
	return "AuditerPlugin"
}

func (p *AuditerPlugin) Initialize(db *gorm.DB) error {
	db.Callback().Create().Before("gorm:create").Register("auditer:before_create", beforeCreateCallback)
	db.Callback().Create().After("gorm:create").Register("auditer:after_create", afterCreateCallback)
	db.Callback().Update().Before("gorm:update").Register("auditer:before_update", beforeUpdateCallback)
	db.Callback().Update().After("gorm:update").Register("auditer:after_update", afterUpdateCallback)
	db.Callback().Delete().Before("gorm:delete").Register("auditer:before_delete", beforeDeleteCallback)
	db.Callback().Delete().After("gorm:delete").Register("auditer:after_delete", afterDeleteCallback)
	return nil
}

func beforeCreateCallback(db *gorm.DB) {
	log.Println("Before create callback")
}

func afterCreateCallback(db *gorm.DB) {
	log.Println("After create callback")
}

func beforeUpdateCallback(db *gorm.DB) {
	log.Println("Before update callback")
}

func afterUpdateCallback(db *gorm.DB) {
	log.Println("After update callback")
}

func beforeDeleteCallback(db *gorm.DB) {
	log.Println("Before delete callback")
}

func afterDeleteCallback(db *gorm.DB) {
	log.Println("After delete callback")
}

