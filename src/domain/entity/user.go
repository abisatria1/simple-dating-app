package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/abisatria1/simple-dating-app/src/constant"
	"github.com/abisatria1/simple-dating-app/src/util"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64              `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name      string             `gorm:"type:varchar(255);not null" json:"name"`
	Email     string             `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password  string             `gorm:"type:varchar(255);not null" json:"-"`
	BirthDate string             `gorm:"type:varchar(20);not null;" json:"birth_date"`
	Gender    constant.Gender    `gorm:"type:smallint(1);not null;" json:"gender"`
	Bio       util.NullString    `gorm:"type:text;" json:"bio"`
	Location  util.NullString    `gorm:"type:text;" json:"location"`
	Religion  *constant.Religion `gorm:"type:smallint(1);" json:"religion"`
	Height    util.NullInt64     `gorm:"type:smallint(1);" json:"height"`
	UserType  constant.UserType  `gorm:"type:smallint(1);not null;default:0" json:"user_type"`
	Quota     int64              `gorm:"type:integer;not null" json:"quota"`
	Photos    Photos             `gorm:"type:json" json:"photos"`
	CreatedAt time.Time          `gorm:"autoCreateTime;type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time          `gorm:"autoUpdateTime;type:timestamp;default:current_timestamp" json:"updated_at"`
	Interests []Interest         `gorm:"many2many:user_interests;" json:"interests"`
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

type Photos []string

func (m *Photos) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, &m)
	case string:
		return json.Unmarshal([]byte(v), &m)
	default:
		return fmt.Errorf("unsupported type for Photos: %T", src)
	}
}

func (m Photos) Value() (driver.Value, error) {
	return json.Marshal(m)
}
