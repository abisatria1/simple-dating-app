package middleware

import (
	"log"
	"time"

	"github.com/abisatria1/simple-dating-app/pkg/custerr"
	"github.com/abisatria1/simple-dating-app/src/constant"
	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/labstack/echo/v4"
)

func (m *Middleware) CheckSubscriptionExpired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		ctx := c.Request().Context()
		user := c.Get(constant.AppUserContext).(*entity.User)
		if user == nil {
			err = custerr.ErrInvalidUser
			return
		}

		subs, err := m.SubscriptionDomain.GetActiveUserSubscription(ctx, user.ID)
		if err != nil {
			log.Printf("[CheckSubscriptionExpired] GetActiveUserSubscription error cause: %s", err.Error())
			return next(c)
		}
		if subs != nil && subs.ExpiredAt.Before(time.Now()) {
			_, err = m.UserDomain.UpdateUser(ctx, user.ID, entity.User{UserType: constant.UserTypeNormal}, nil)
			if err != nil {
				log.Printf("[CheckSubscriptionExpired] update user subscription type failed cause: %s", err.Error())
			}
		}

		return next(c)
	}
}
