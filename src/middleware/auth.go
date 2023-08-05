package middleware

import (
	"net/http"
	"strings"

	"github.com/abisatria1/simple-dating-app/pkg/jwt"
	"github.com/abisatria1/simple-dating-app/src/constant"
	"github.com/abisatria1/simple-dating-app/src/domain/subscription"
	"github.com/abisatria1/simple-dating-app/src/domain/user"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	UserDomain         user.UserDomainManager
	SubscriptionDomain subscription.SubscriptionDomainManager
	Jwt                jwt.JwtManager
}

type Options struct {
	UserDomain         user.UserDomainManager
	JwtManager         jwt.JwtManager
	SubscriptionDomain subscription.SubscriptionDomainManager
}

func New(o *Options) Middleware {
	return Middleware{
		UserDomain:         o.UserDomain,
		Jwt:                o.JwtManager,
		SubscriptionDomain: o.SubscriptionDomain,
	}
}

func (m *Middleware) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
		}

		jwtToken := ""
		if len(tokenString) > 7 && strings.ToUpper(tokenString[0:7]) == "BEARER " {
			jwtToken = tokenString[7:]
		}
		if jwtToken == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token format")
		}

		claims, err := m.Jwt.ParseToken(jwtToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token").SetInternal(err)
		}
		if claims == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "jwt payload is empty").SetInternal(err)
		}

		user, err := m.UserDomain.GetUserByID(c.Request().Context(), claims.UserID)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "user not found")
		}
		c.Set(constant.AppUserContext, user)

		return next(c)
	}
}
