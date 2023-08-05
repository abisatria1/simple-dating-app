package custerr

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrInvalidPayload                      = echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid request payload")
	ErrInvalidUser                         = echo.NewHTTPError(http.StatusUnauthorized, "no user was found")
	ErrRecordNotFound                      = echo.NewHTTPError(http.StatusBadRequest, "record not found")
	ErrNoQuota                             = echo.NewHTTPError(http.StatusBadRequest, "no quota left")
	ErrSameAccount                         = echo.NewHTTPError(http.StatusBadRequest, "same account can't do this action")
	ErrUserAlreadyAppearInFeeds            = echo.NewHTTPError(http.StatusBadRequest, "same account appear more than once, can't do this action")
	ErrSameSubscription                    = echo.NewHTTPError(http.StatusBadRequest, "user subscription already same")
	ErrCantUpgradeIntoSelectedSubscription = echo.NewHTTPError(http.StatusBadRequest, "can't upgrade subscription")
)
