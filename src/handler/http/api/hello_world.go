package api

import (
	"github.com/abisatria1/simple-dating-app/pkg/httpservice"
	"github.com/labstack/echo/v4"
)

func (a *API) Hello(c echo.Context) (response httpservice.HandlerResponse, err error) {
	return httpservice.NewJsonResponse().SetData(map[string]interface{}{"test": "mantap"}), nil
}
