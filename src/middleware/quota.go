package middleware

import (
	"github.com/abisatria1/simple-dating-app/pkg/custerr"
	"github.com/abisatria1/simple-dating-app/src/constant"
	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/labstack/echo/v4"
)

func (m *Middleware) QuotaCheckingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		user := c.Get(constant.AppUserContext).(*entity.User)
		if user == nil {
			err = custerr.ErrInvalidUser
			return
		}

		if !m.UserDomain.ValidateUserQuota(c.Request().Context(), *user) {
			err = custerr.ErrNoQuota
			return
		}

		return next(c)
	}
}
