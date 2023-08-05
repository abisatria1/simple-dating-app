package entity

import "time"

type Interest struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name      string    `gorm:"type:varchar(50);not null;index" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime;type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;type:timestamp;default:current_timestamp" json:"updated_at"`
}
