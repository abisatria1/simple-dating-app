package model

import (
	"time"

	"github.com/abisatria1/simple-dating-app/src/constant"
	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/asaskevich/govalidator"
)

type UpgradeSubscriptionRequest struct {
	Duration         int64             `json:"duration" valid:"required"`
	SubscriptionType constant.UserType `json:"subscription_type" valid:"required"`
}

func (model *UpgradeSubscriptionRequest) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(model); err != nil {
		return
	}
	return
}

func (model *UpgradeSubscriptionRequest) ToModel() (subs entity.Subscription, err error) {
	subs = entity.Subscription{
		UserType:  model.SubscriptionType,
		Duration:  model.Duration,
		ExpiredAt: time.Now().Add(time.Duration(model.Duration*24) * time.Hour),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return
}
