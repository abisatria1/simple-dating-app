package auth

import (
	"context"
	"net/http"

	"github.com/abisatria1/simple-dating-app/src/model"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (m *authUseCase) RegisterUser(ctx context.Context, payload model.RegisterUserRequest) (ID int64, err error) {
	user := model.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	if err = user.HashPassword(); err != nil {
		err = errors.Wrapf(err, "hashed password err: %s", err.Error())
		return
	}

	checkEmail, err := m.UserDomain.GetUserByEmail(ctx, user.Email)
	if err != nil {
		err = errors.Wrapf(err, "GetUserByEmail error : %s", err.Error())
		return
	}
	if checkEmail != nil {
		err = echo.NewHTTPError(http.StatusBadRequest, "email already used")
		return
	}

	if ID, err = m.UserDomain.CreateUser(ctx, user, nil); err != nil {
		err = errors.Wrapf(err, "CreateUser error : %s", err.Error())
		return
	}

	return
}
