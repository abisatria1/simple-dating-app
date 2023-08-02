package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime;type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;type:timestamp;default:current_timestamp" json:"updated_at"`
}

func (u *User) HashPassword() (err error) {
	password := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return
	}
	u.Password = string(hashedPassword)
	return
}

func (u *User) ComparePassword(nonHashPassword string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(nonHashPassword))
}
