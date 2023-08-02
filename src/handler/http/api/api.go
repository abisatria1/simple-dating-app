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
}
