package model

import (
	"github.com/asaskevich/govalidator"
)

type LoginRequest struct {
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
}

func (model *LoginRequest) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(model); err != nil {
		return
	}
	return
}
