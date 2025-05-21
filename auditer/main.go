package auditer

import (
	"gorm.io/gorm"
)

var GlobalAuditLog AuditLog

func InitSDK(db *gorm.DB, user User) {
	GlobalAuditLog = AuditLog{
		Service: "test",
		User:    user,
	}
	db.Use(&AuditerPlugin{})
}
