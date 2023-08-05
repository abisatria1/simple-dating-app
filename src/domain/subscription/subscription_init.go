package subscription

import (
	"context"
	"time"

	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SubscriptionDomainManager interface {
	Begin(ctx context.Context) (tx *gorm.DB)
	InsertSubscription(ctx context.Context, subs entity.Subscription, tx *gorm.DB) (ID int64, err error)
	GetActiveUserSubscription(ctx context.Context, userID int64) (subs *entity.Subscription, err error)
}

type SubscriptionDomain struct {
	DB *gorm.DB
}

type Options struct {
	DB *gorm.DB
}

func New(o *Options) SubscriptionDomainManager {
	return &SubscriptionDomain{
		DB: o.DB,
	}
}

func (u *SubscriptionDomain) Begin(ctx context.Context) (tx *gorm.DB) {
	tx = u.DB.Begin()
	return
}

func (u *SubscriptionDomain) InsertSubscription(ctx context.Context, subs entity.Subscription, tx *gorm.DB) (ID int64, err error) {
	executor := u.DB
	if tx != nil {
		executor = tx
	}
	query := executor.WithContext(ctx).Create(&subs)
	if err = query.Error; err != nil {
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	ID = subs.ID
	return
}

func (u *SubscriptionDomain) GetActiveUserSubscription(ctx context.Context, userID int64) (subs *entity.Subscription, err error) {
	subs = &entity.Subscription{}
	query := u.DB.WithContext(ctx).Where("user_id = ? and expired_at >= ?", userID, time.Now()).Order("expired_at desc").First(subs)
	if err = query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	return
}
