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

func (a *API) GetUserFeeds(c echo.Context) (response httpservice.HandlerResponse, err error) {
	response = httpservice.NewJsonResponse()
	user := c.Get(constant.AppUserContext).(*entity.User)
	query := model.GetFeedsForUserParam{}

	if user == nil {
		err = custerr.ErrInvalidUser
		return
	}
	if err = c.Bind(&query); err != nil {
		err = custerr.ErrInvalidPayload.SetInternal(errors.Wrapf(err, "error bind: %s", err.Error()))
		return
	}
	result, err := a.service.Usecases.Feeds.GetFeedsForUser(c.Request().Context(), *user, query)
	if err != nil {
		return
	}

	return httpservice.NewJsonResponse().SetData(result), nil
}

func (a *API) Dislike(c echo.Context) (response httpservice.HandlerResponse, err error) {
	response = httpservice.NewJsonResponse()
	user := c.Get(constant.AppUserContext).(*entity.User)
	body := model.DislikeRequest{}

	if user == nil {
		err = custerr.ErrInvalidUser
		return
	}
	if err = c.Bind(&body); err != nil {
		err = custerr.ErrInvalidPayload.SetInternal(errors.Wrapf(err, "error bind: %s", err.Error()))
		return
	}

	if err = a.service.Usecases.Feeds.Dislike(c.Request().Context(), *user, body.TargetUserID); err != nil {
		return
	}

	return httpservice.NewJsonResponse().SetMessage("sucess"), nil
}

func (a *API) Like(c echo.Context) (response httpservice.HandlerResponse, err error) {
	response = httpservice.NewJsonResponse()
	user := c.Get(constant.AppUserContext).(*entity.User)
	body := model.DislikeRequest{}

	if user == nil {
		err = custerr.ErrInvalidUser
		return
	}
	if err = c.Bind(&body); err != nil {
		err = custerr.ErrInvalidPayload.SetInternal(errors.Wrapf(err, "error bind: %s", err.Error()))
		return
	}

	if err = a.service.Usecases.Feeds.Like(c.Request().Context(), *user, body.TargetUserID); err != nil {
		return
	}

	return httpservice.NewJsonResponse().SetMessage("sucess"), nil
}
