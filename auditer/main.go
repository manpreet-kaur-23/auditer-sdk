package auditer

import (
	"gorm.io/gorm"

	"github.com/manpreet-kaur-23/auditer-sdk/constants"
)

var GlobalAuditLog constants.AuditLog

func InitSDK(db *gorm.DB, user constants.User) {
	GlobalAuditLog = constants.AuditLog{
		Service: "test",
		User:    user,
	}
	db.Use(&AuditerPlugin{})
}
