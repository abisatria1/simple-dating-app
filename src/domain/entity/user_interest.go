package entity

import "time"

type UserInterest struct {
	ID         int64     `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	UserID     int64     `gorm:"type:integer;not null;index:idx_user_interest,unique" json:"user_id"`
	InterestID int64     `gorm:"type:integer;not null;index:idx_user_interest,unique" json:"interest_id"`
	CreatedAt  time.Time `gorm:"autoCreateTime;type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime;type:timestamp;default:current_timestamp" json:"updated_at"`
}
