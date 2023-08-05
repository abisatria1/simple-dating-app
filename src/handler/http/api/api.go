package api

import (
	"github.com/abisatria1/simple-dating-app/pkg/httpservice"
	"github.com/abisatria1/simple-dating-app/src/service"
	"github.com/labstack/echo/v4"
)

type Options struct {
	Prefix  string
	Service *service.Service
}

type API struct {
	prefix  string
	service *service.Service
}

func New(o *Options) *API {
	return &API{
		prefix:  o.Prefix,
		service: o.Service,
	}
}

func (a *API) Register(e *echo.Echo) {
	r := e.Group(a.prefix)

	r.GET("/test", httpservice.DefaultHandler(a.Hello))
	r.POST("/login", httpservice.DefaultHandler(a.Login))
	r.POST("/register", httpservice.DefaultHandler(a.RegisterUser))

	// interest
	r.GET("/interest", httpservice.DefaultHandler(a.GetAllInterests))

	// user
	r.Use(a.service.Middleware.JWTMiddleware, a.service.Middleware.CheckSubscriptionExpired)
	r.GET("/user", httpservice.DefaultHandler(a.GetLoggedUser))
	r.POST("/user/interest", httpservice.DefaultHandler(a.InsertUserInterest))
	r.GET("/user/feeds", httpservice.DefaultHandler(a.GetUserFeeds), a.service.Middleware.QuotaCheckingMiddleware)
	r.POST("/user/feeds/dislike", httpservice.DefaultHandler(a.Dislike), a.service.Middleware.QuotaCheckingMiddleware)
	r.POST("/user/feeds/like", httpservice.DefaultHandler(a.Like), a.service.Middleware.QuotaCheckingMiddleware)
	r.GET("/user/match", httpservice.DefaultHandler(a.GetUserMatchList))
	r.GET("/user/match/potential", httpservice.DefaultHandler(a.GetUserPotentialMatchList))

	// user subscription
	r.POST("/user/subscription/upgrade", httpservice.DefaultHandler(a.UpgradeSubscription))
	r.GET("/user/subscription", httpservice.DefaultHandler(a.GetActiveSubscription))
}
