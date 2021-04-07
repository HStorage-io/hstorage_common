package hstorage_common

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type BlockUser struct {
	gorm.Model
	Attributes datatypes.JSON
}

type BlockUserAttributes struct {
	IPAddress string   `json:"ip_address"`
	UserIDs   []string `json:"user_ids"`
	Reason    string   `json:"reason"`
}

type UserSetting struct {
	gorm.Model                 `json:"-"`
	UserID                     string `gorm:"type:varchar(255);uniqueIndex" json:"user_id"`
	EnableAutoDelete           bool   `gorm:"type:tinyint(1);default:0" json:"enable_auto_delete"`
	AutomaticDeletionSeconds   uint   `gorm:"type:int unsigned" json:"automatic_deletion_seconds"`
	EnableAutoPassword         bool   `gorm:"type:tinyint(1);default:0" json:"enable_auto_password"`
	Password                   string `gorm:"type:varchar(255)" json:"password"`
	EnableAutoEncryption       bool   `gorm:"type:tinyint(1);default:0" json:"enable_auto_encryption"`
	EnableDownloadNotification bool   `gorm:"type:tinyint(1);default:0" json:"enable_download_notification"`
}

type SFTPUser struct {
	gorm.Model
	UserID     string `gorm:"type:varchar(255);uniqueIndex" json:"user_id"`
	PrivateKey string `gorm:"type:text(500);not null" json:"private_key"`
	PublicKey  string `gorm:"type:varchar(100);not null" json:"public_key"`
}
