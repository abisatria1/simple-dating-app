package subscription

import (
	"context"

	"github.com/abisatria1/simple-dating-app/pkg/custerr"
	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/model"
	"github.com/pkg/errors"
)

func (m *subscriptionUseCase) UpgradeSubscription(ctx context.Context, user entity.User, subscription model.UpgradeSubscriptionRequest) (subsID int64, err error) {
	if user.UserType == subscription.SubscriptionType {
		err = custerr.ErrSameSubscription
		return
	}

	if !subscription.SubscriptionType.IsEligibleUpgrade() {
		err = custerr.ErrCantUpgradeIntoSelectedSubscription
		return
	}

	subs, _ := subscription.ToModel()
	subs.UserID = user.ID
	tx := m.SubscriptionDomain.Begin(ctx)
	defer tx.Rollback()
	if subsID, err = m.SubscriptionDomain.InsertSubscription(ctx, subs, tx); err != nil {
		err = errors.Wrapf(err, "[UpgradeSubscription] InsertSubscription error: %s", err)
		return
	}

	if _, err = m.UserDomain.UpdateUser(ctx, user.ID, entity.User{UserType: subs.UserType}, tx); err != nil {
		err = errors.Wrapf(err, "[UpgradeSubscription] UpdateUser error: %s", err)
		return
	}

	tx.Commit()
	return
}

func (m *subscriptionUseCase) GetActiveSubscription(ctx context.Context, user entity.User) (subs *entity.Subscription, err error) {
	if subs, err = m.SubscriptionDomain.GetActiveUserSubscription(ctx, user.ID); err != nil {
		err = errors.Wrapf(err, "[GetActiveSubscription] GetActiveUserSubscription error: %s", err)
		return
	}
	return
}
