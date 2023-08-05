package entity

import "time"

type UserBlacklist struct {
	ID           int64     `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	UserID       int64     `gorm:"type:integer;not null;index:idx_user_blacklist" json:"user_id"`
	TargetUserID int64     `gorm:"type:integer;not null;index:idx_user_blacklist" json:"target_user_id"`
	ExpiredAt    time.Time `gorm:"type:timestamp;not null;index:idx_user_blacklist" json:"expired_at"`
	CreatedAt    time.Time `gorm:"autoCreateTime;type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;type:timestamp;default:current_timestamp" json:"updated_at"`
}
