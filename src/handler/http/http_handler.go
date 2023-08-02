package http

import (
	"github.com/abisatria1/simple-dating-app/src/handler/http/api"
	"github.com/abisatria1/simple-dating-app/src/service"
	"github.com/labstack/echo/v4"
)

type HttpHandler struct {
	handlers []Handler
}

type Handler interface {
	Register(e *echo.Echo)
}

func New(service *service.Service) *HttpHandler {
	handlers := []Handler{
		api.New(&api.Options{
			Prefix:  "/api",
			Service: service,
		}),
	}

	return &HttpHandler{
		handlers: handlers,
	}
}

func (h *HttpHandler) RegisterHandlers(e *echo.Echo) {
	for i := range h.handlers {
		h.handlers[i].Register(e)
	}
}
