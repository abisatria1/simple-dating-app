package handler

import (
	"github.com/abisatria1/simple-dating-app/src/handler/http"
	"github.com/abisatria1/simple-dating-app/src/service"
	"github.com/labstack/echo/v4"
)

type HandlerManager interface {
	RegisterHttpHandlers(echo *echo.Echo)
}

type Handler struct {
	httpHandler *http.HttpHandler
}

func (h *Handler) RegisterHttpHandlers(echo *echo.Echo) {
	h.httpHandler.RegisterHandlers(echo)
}

func New(service *service.Service) HandlerManager {
	return &Handler{
		httpHandler: http.New(service),
	}
}
