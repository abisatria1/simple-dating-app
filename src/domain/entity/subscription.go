package entity

import (
	"time"

	"github.com/abisatria1/simple-dating-app/src/constant"
)

type Subscription struct {
	ID        int64             `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	UserID    int64             `gorm:"type:integer;not null;index" json:"user_id"`
	UserType  constant.UserType `gorm:"type:smallint(1);not null" json:"subscription_type"`
	Duration  int64             `gorm:"type:smallint(1);not null;default:0" json:"duration"`
	ExpiredAt time.Time         `gorm:"type:timestamp;not null" json:"expired_at"`
	CreatedAt time.Time         `gorm:"autoCreateTime;type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time         `gorm:"autoUpdateTime;type:timestamp;default:current_timestamp" json:"updated_at"`
}
