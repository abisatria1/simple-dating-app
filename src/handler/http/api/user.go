package api

import (
	"github.com/abisatria1/simple-dating-app/pkg/custerr"
	"github.com/abisatria1/simple-dating-app/pkg/httpservice"
	"github.com/abisatria1/simple-dating-app/src/constant"
	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/model"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (a *API) RegisterUser(c echo.Context) (response httpservice.HandlerResponse, err error) {
	payload := model.RegisterUserRequest{}
	response = httpservice.NewJsonResponse()

	if err = c.Bind(&payload); err != nil {
		err = custerr.ErrInvalidPayload.SetInternal(errors.Wrapf(err, "error bind: %s", err.Error()))
		return
	}

	if err = payload.Validate(); err != nil {
		err = custerr.ErrInvalidPayload.SetInternal(errors.Wrapf(err, "validation err: %s", err.Error()))
		return
	}

	userId, err := a.service.Usecases.User.RegisterUser(c.Request().Context(), payload)
	if err != nil {
		return
	}

	return httpservice.NewJsonResponse().SetData(map[string]interface{}{"id": userId}), nil
}

func (a *API) GetLoggedUser(c echo.Context) (response httpservice.HandlerResponse, err error) {
	response = httpservice.NewJsonResponse()
	user := c.Get(constant.AppUserContext).(*entity.User)
	return httpservice.NewJsonResponse().SetData(user), nil
}

func (a *API) GetUserPotentialMatchList(c echo.Context) (response httpservice.HandlerResponse, err error) {
	response = httpservice.NewJsonResponse()
	user := c.Get(constant.AppUserContext).(*entity.User)

	if user == nil {
		err = custerr.ErrInvalidUser
		return
	}
	result, err := a.service.Usecases.User.GetUserPotentialMatchList(c.Request().Context(), *user)
	if err != nil {
		return
	}

	return httpservice.NewJsonResponse().SetData(result), nil
}

func (a *API) GetUserMatchList(c echo.Context) (response httpservice.HandlerResponse, err error) {
	response = httpservice.NewJsonResponse()
	user := c.Get(constant.AppUserContext).(*entity.User)

	if user == nil {
		err = custerr.ErrInvalidUser
		return
	}
	result, err := a.service.Usecases.User.GetUserMatchList(c.Request().Context(), *user)
	if err != nil {
		return
	}

	return httpservice.NewJsonResponse().SetData(result), nil
}
