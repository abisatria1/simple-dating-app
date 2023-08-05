package subscription

import (
	"context"

	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/domain/subscription"
	"github.com/abisatria1/simple-dating-app/src/domain/user"
	"github.com/abisatria1/simple-dating-app/src/model"
)

type SubscriptionUseCaseManager interface {
	UpgradeSubscription(ctx context.Context, user entity.User, subscription model.UpgradeSubscriptionRequest) (subsID int64, err error)
	GetActiveSubscription(ctx context.Context, user entity.User) (subs *entity.Subscription, err error)
}

type Options struct {
	UserDomain         user.UserDomainManager
	SubscriptionDomain subscription.SubscriptionDomainManager
}

type subscriptionUseCase struct {
	UserDomain         user.UserDomainManager
	SubscriptionDomain subscription.SubscriptionDomainManager
}

func New(o *Options) SubscriptionUseCaseManager {
	return &subscriptionUseCase{
		UserDomain:         o.UserDomain,
		SubscriptionDomain: o.SubscriptionDomain,
	}
}
