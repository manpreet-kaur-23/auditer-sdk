package auditer

import (
	"gorm.io/gorm"
)

func InitSDK(db *gorm.DB, user User) {
	db.Use(&AuditerPlugin{})
}
