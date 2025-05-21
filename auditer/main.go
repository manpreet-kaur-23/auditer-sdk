package auditer

import (
	"time"

	"gorm.io/gorm"

	"github.com/manpreet-kaur-23/auditer-sdk/constants"
)

var GlobalAuditLog constants.AuditLog

func InitSDK(db *gorm.DB) {
	db.Use(&AuditerPlugin{})
}

func InitializeAuditLog(user constants.User, service string) {
	GlobalAuditLog.User = user
	GlobalAuditLog.Timestamp = time.Now()
	GlobalAuditLog.Service = service
}
