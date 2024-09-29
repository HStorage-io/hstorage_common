package hstorage_common

import "gorm.io/gorm"

type DeveloperKey struct {
	gorm.Model
	UserID    string `gorm:"type:varchar(255); not null; index:idx_user_id_developer" json:"user_id"`
	SecretKey string `gorm:"type:varchar(255); not null; index:idx_secret_key_developer" json:"secret_key"`
}
