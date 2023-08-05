package model

import (
	"database/sql"
	"time"

	"github.com/abisatria1/simple-dating-app/src/constant"
	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/util"
	"github.com/asaskevich/govalidator"
)

type RegisterUserRequest struct {
	Name      string             `json:"name" valid:"required"`
	Email     string             `json:"email" valid:"email,required"`
	Password  string             `json:"password" valid:"required"`
	BirthDate string             `json:"birth_date" valid:"required"`
	Gender    constant.Gender    `json:"gender" valid:"required"`
	Photos    []string           `json:"photos" valid:"required"`
	Bio       *string            `json:"bio"`
	Location  *string            `json:"location"`
	Religion  *constant.Religion `json:"religion"`
	Height    *int64             `json:"height"`
}

func (model *RegisterUserRequest) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(model); err != nil {
		return
	}
	return
}

func (model *RegisterUserRequest) ToModel() (user entity.User, err error) {
	user = entity.User{
		Name:      model.Name,
		Email:     model.Email,
		Password:  model.Password,
		BirthDate: model.BirthDate,
		Gender:    model.Gender,
		Quota:     constant.DefaultNormalUserQuota,
		UserType:  constant.UserTypeNormal,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if model.Bio != nil {
		user.Bio = util.NullString{NullString: sql.NullString{String: *model.Bio, Valid: true}}
	}
	if model.Location != nil {
		user.Location = util.NullString{NullString: sql.NullString{String: *model.Location, Valid: true}}
	}
	if model.Religion != nil {
		user.Religion = model.Religion
	}
	if model.Height != nil {
		user.Height = util.NullInt64{NullInt64: sql.NullInt64{Int64: *model.Height, Valid: true}}
	}

	user.Photos = model.Photos

	return
}

type UserFeed struct {
	ID                int64              `json:"id"`
	Name              string             `json:"name"`
	BirthDate         string             `json:"birth_date"`
	Gender            constant.Gender    `json:"gender"`
	Bio               util.NullString    `json:"bio"`
	Location          util.NullString    `json:"location"`
	Religion          *constant.Religion `json:"religion"`
	Height            util.NullInt64     `json:"height"`
	IsProfileVerified bool               `json:"is_profile_verified"`
	Photos            entity.Photos      `json:"photos"`
	Interests         []entity.Interest  `json:"interests"`
}
