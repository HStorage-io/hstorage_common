package hstorage_common

import "gorm.io/gorm"

type PlanType string

const (
	PlanTypeFree     PlanType = "free"
	PlanTypePremium  PlanType = "premium"
	PlanTypeBusiness PlanType = "business"
	PlanTypeAPI      PlanType = "api"
)

type SubscriptionFailed struct {
	gorm.Model
	UserID         string   `gorm:"varchar(255);not null"`
	CustomerID     string   `gorm:"varchar(255);not null"`
	Email          string   `gorm:"varchar(255);not null"`
	PlanType       PlanType `gorm:"varchar(255);not null"`
	ReadyForDelete bool     `gorm:"default:false;comment:ユーザーがサイトから退会処理をした場合に true になる。false の場合、CreatedAt から n日後に実行される"`
}
