package httpservice

import (
	"time"

	"github.com/labstack/echo/v4"
)

type HandlerResponse interface {
	Send(c echo.Context) error
	SetLatency(float64) *JSONResponse
	SetMessage(string) *JSONResponse
}

func DefaultHandler(f func(echo.Context) (HandlerResponse, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		now := time.Now()
		response, err := f(c)
		if err != nil {
			return err
		}
		response.SetLatency(float64(time.Since(now).Milliseconds()))
		return response.Send(c)
	}
}
