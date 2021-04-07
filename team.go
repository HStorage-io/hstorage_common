package hstorage_common

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	LeaderUserID    string `gorm:"type:varchar(255);not null" json:"leader_user_id"`
	LeaderUserEmail string `gorm:"type:varchar(255);not null" json:"leader_user_email"`
	MemberUserID    string `gorm:"type:varchar(255);not null;uniqueIndex" json:"member_user_id"`
	MemberUserEmail string `gorm:"type:varchar(255);not null;uniqueIndex" json:"member_user_email"`
	Status          int    `gorm:"type:tinyint(3);not null;default:0;comment:0:pending 1:billed 2:delete" json:"status"`
}
