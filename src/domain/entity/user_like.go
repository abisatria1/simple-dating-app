package entity

import "time"

type UserLike struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	LikeBy    int64     `gorm:"type:integer;not null;index" json:"like_by"`
	Like      int64     `gorm:"type:integer;not null;index" json:"like"`
	ExpiredAt time.Time `gorm:"type:timestamp;not null;index" json:"expired_at"`
	CreatedAt time.Time `gorm:"autoCreateTime;type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;type:timestamp;default:current_timestamp" json:"updated_at"`
}
