package echohttp

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Options struct {
	ListenAddress int
	WriteTimeout  int
	ReadTimeout   int
}

type EchoHttpServer interface {
	Run()
	GetEcho() *echo.Echo
	ListenError() chan error
}

type httpServer struct {
	writeTimeout  int
	readTimeout   int
	listenAddress int
	echo          *echo.Echo
	errorCh       chan error
}

func New(o *Options) EchoHttpServer {
	return &httpServer{
		writeTimeout:  o.WriteTimeout,
		readTimeout:   o.ReadTimeout,
		listenAddress: o.ListenAddress,
		echo:          echo.New(),
	}
}

func (h *httpServer) Run() {
	h.echo.Use(middleware.Logger())
	h.echo.Use(middleware.Recover())
	h.echo.Use(middleware.Secure())
	h.echo.Use(middleware.CORS())
	h.echo.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Duration(h.readTimeout) * time.Millisecond,
	}))

	fmt.Printf("starting application on %d", h.listenAddress)
	// h.echo.HTTPErrorHandler = handleEchoError()
	if err := h.echo.Start(fmt.Sprintf(":%d", h.listenAddress)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		h.errorCh <- err
	}

}

func (h *httpServer) GetEcho() *echo.Echo {
	return h.echo
}

func (h *httpServer) ListenError() chan error {
	return h.errorCh
}

func handleEchoError() echo.HTTPErrorHandler {
	return func(err error, ctx echo.Context) {
		// var echoError *echo.HTTPError

		// // if *echo.HTTPError, let echokit middleware handles it
		// if errors.As(err, &echoError) {
		// 	return
		// }

		// statusCode := http.StatusInternalServerError
		// // message := "mohon maaf, terjadi kesalahan pada server"
		// message := err.Error()

		// switch {
		// case errors.Is(err, httpservice.ErrBadRequest) || errors.Is(err, httpservice.ErrPasswordNotMatch) || errors.Is(err, httpservice.ErrConfirmPasswordNotMatch):
		// 	statusCode = http.StatusBadRequest
		// 	message = err.Error()
		// case errors.Is(err, httpservice.ErrInvalidAppKey) || errors.Is(err, httpservice.ErrInvalidOTP) || errors.Is(err, httpservice.ErrUnauthorizedUser) || errors.Is(err, httpservice.ErrInActiveUser) || errors.Is(err, httpservice.ErrUnauthorizedTokenData):
		// 	statusCode = http.StatusUnauthorized
		// 	message = err.Error()
		// case errors.Is(err, httpservice.ErrUserNotFound):
		// 	statusCode = http.StatusNotFound
		// 	message = err.Error()
		// case errors.Is(err, httpservice.ErrNoResultData):
		// 	statusCode = http.StatusOK
		// 	message = err.Error()
		// }

		// _ = ctx.JSON(statusCode, echo.NewHTTPError(statusCode, message))
	}
}
