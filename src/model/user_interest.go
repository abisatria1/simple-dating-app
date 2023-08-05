package model

import (
	"github.com/asaskevich/govalidator"
)

type InsertUserInterestRequest struct {
	InterestIds []int64 `json:"interest_ids" valid:"required"`
}

type InsertUserInterestRequestItem struct {
}

func (m *InsertUserInterestRequest) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(m); err != nil {
		return
	}
	return
}
