package hstorage_common

type Group struct {
	ID             uint     `gorm:"primaryKey" json:"id"`
	UID            string   `gorm:"not null;unique;size:32" json:"uid"`
	UserID         string   `gorm:"type:varchar(255);not null;index" json:"-"`
	Name           string   `gorm:"type:varchar(255);not null" json:"name"`
	IsPublicView   bool     `gorm:"type:tinyint(2);default:0" json:"is_public_view"`
	IsPublicUpload bool     `gorm:"type:tinyint(2);default:0" json:"is_public_upload"`
	Uploads        []Upload `gorm:"-" json:"uploads"`
}
