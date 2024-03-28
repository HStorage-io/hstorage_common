package hstorage_common

type Setting struct {
	ID                        int    `gorm:"primaryKey" json:"-"`
	MaxUploadSize             uint64 `gorm:"comment: maximum upload size" json:"max_upload_size"`
	FreeUserMonthlyCountLimit uint8  `gorm:"comment: 月間のアップロード数" json:"free_user_monthly_count_limit"`
	FreeUserCapacityLimit     uint64 `gorm:"comment: アカウント単位での最大保管容量 byte" json:"free_user_capacity_limit"`
	FreeUserCountLimit        int64  `gorm:"comment: アカウント単位での最大保管ファイル数" json:"free_user_count_limit"`
	PremiumLimitUploadSize    uint64 `gorm:"comment: byte" json:"premium_limit_upload_size"`
	PremiumMaxFiles           uint64 `gorm:"comment: dropzone maxFiles" json:"premium_max_files"`
	BusinessLimitUploadSize   uint64 `gorm:"comment: byte" json:"business_limit_upload_size"`
	BusinessMaxFiles          uint64 `gorm:"comment: dropzone maxFiles" json:"business_max_files"`
}
