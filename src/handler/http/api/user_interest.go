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

func (a *API) GetAllInterests(c echo.Context) (response httpservice.HandlerResponse, err error) {
	response = httpservice.NewJsonResponse()
	result, err := a.service.Usecases.User.GetAllInterest(c.Request().Context())
	if err != nil {
		return
	}
	return httpservice.NewJsonResponse().SetData(result), nil
}

func (a *API) InsertUserInterest(c echo.Context) (response httpservice.HandlerResponse, err error) {
	payload := model.InsertUserInterestRequest{}
	response = httpservice.NewJsonResponse()
	user := c.Get(constant.AppUserContext).(*entity.User)

	if user == nil {
		err = custerr.ErrInvalidUser
		return
	}

	if err = c.Bind(&payload); err != nil {
		err = custerr.ErrInvalidPayload.SetInternal(errors.Wrapf(err, "error bind: %s", err.Error()))
		return
	}

	if err = payload.Validate(); err != nil {
		err = custerr.ErrInvalidPayload.SetInternal(errors.Wrapf(err, "validation err: %s", err.Error()))
		return
	}

	insertedRows, err := a.service.Usecases.User.InsertUserInterest(c.Request().Context(), user.ID, payload)
	if err != nil {
		return
	}

	return httpservice.NewJsonResponse().SetData(map[string]interface{}{"inserted_rows": insertedRows}), nil
}

func (a *API) GetUserInterest(c echo.Context) (response httpservice.HandlerResponse, err error) {
	// payload := model.RegisterUserRequest{}
	// response = httpservice.NewJsonResponse()

	// if err = c.Bind(&payload); err != nil {
	// 	err = custerr.ErrInvalidPayload.SetInternal(errors.Wrapf(err, "error bind: %s", err.Error()))
	// 	return
	// }

	// if err = payload.Validate(); err != nil {
	// 	err = custerr.ErrInvalidPayload.SetInternal(errors.Wrapf(err, "validation err: %s", err.Error()))
	// 	return
	// }

	// userId, err := a.service.Usecases.Auth.RegisterUser(c.Request().Context(), payload)
	// if err != nil {
	// 	return
	// }

	// return httpservice.NewJsonResponse().SetData(map[string]interface{}{"id": userId}), nil
	return
}
