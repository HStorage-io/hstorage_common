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

func ErrGroupIDNotProvided() error {
	return &MyError{
		Err: "group_id not provided",
	}
}

func ErrGroupUIDNotProvided() error {
	return &MyError{
		Err: "group_uid not provided",
	}
}

func ErrGroupTypeNotProvided() error {
	return &MyError{
		Err: "?type is not provided",
	}
}

func ErrGroupNotFound() error {
	return &MyError{
		Err: "group not provided",
	}
}

func ErrGroupNotPublicView() error {
	return &MyError{
		Err: "group_not_public_view",
	}
}

func ErrGroupNotPublicUpload() error {
	return &MyError{
		Err: "group_not_public_upload",
	}
}
