package custerr

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrInvalidPayload = echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid request payload")
	ErrRecordNotFound = echo.NewHTTPError(http.StatusBadRequest, "record not found")
)
