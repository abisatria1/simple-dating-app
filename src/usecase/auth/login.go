package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abisatria1/simple-dating-app/src/model"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (m *authUseCase) Login(ctx context.Context, payload model.LoginRequest) (jwtToken string, err error) {
	user, err := m.UserDomain.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		err = errors.Wrapf(err, "error GetUserByEmail: %s", err.Error())
		return
	}
	if user == nil {
		err = echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user with email %s not found", payload.Email))
		return
	}
	if err = user.ComparePassword(payload.Password); err != nil {
		err = echo.NewHTTPError(http.StatusBadRequest, "invalid password")
		return
	}
	jwtToken = "correct jwt"
	return
}
