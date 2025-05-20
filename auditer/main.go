package auditer

import (
	"gorm.io/gorm"
)

func InitSDK(db *gorm.DB) {

	db.Use(&AuditerPlugin{})
}
