package api

import (
	"github.com/abisatria1/simple-dating-app/pkg/custerr"
	"github.com/abisatria1/simple-dating-app/pkg/httpservice"
	"github.com/abisatria1/simple-dating-app/src/model"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (a *API) Login(c echo.Context) (response httpservice.HandlerResponse, err error) {
	payload := model.LoginRequest{}
	response = httpservice.NewJsonResponse()

	if err = c.Bind(&payload); err != nil {
		err = custerr.ErrInvalidPayload.SetInternal(errors.Wrapf(err, "error bind: %s", err.Error()))
		return
	}

	if err = payload.Validate(); err != nil {
		err = custerr.ErrInvalidPayload.SetInternal(errors.Wrapf(err, "validation err: %s", err.Error()))
		return
	}

	jwtToken, err := a.service.Usecases.Auth.Login(c.Request().Context(), payload)
	if err != nil {
		return
	}

	return httpservice.NewJsonResponse().SetData(jwtToken), nil
}
