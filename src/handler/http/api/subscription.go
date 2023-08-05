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

func (a *API) UpgradeSubscription(c echo.Context) (response httpservice.HandlerResponse, err error) {
	payload := model.UpgradeSubscriptionRequest{}
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

	subsId, err := a.service.Usecases.Subscription.UpgradeSubscription(c.Request().Context(), *user, payload)
	if err != nil {
		return
	}

	return httpservice.NewJsonResponse().SetData(map[string]interface{}{"subscription_id": subsId}), nil
}

func (a *API) GetActiveSubscription(c echo.Context) (response httpservice.HandlerResponse, err error) {
	response = httpservice.NewJsonResponse()
	user := c.Get(constant.AppUserContext).(*entity.User)

	if user == nil {
		err = custerr.ErrInvalidUser
		return
	}

	result, err := a.service.Usecases.Subscription.GetActiveSubscription(c.Request().Context(), *user)
	if err != nil {
		return
	}

	return httpservice.NewJsonResponse().SetData(result), nil
}
