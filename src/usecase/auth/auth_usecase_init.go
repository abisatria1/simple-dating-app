package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abisatria1/simple-dating-app/pkg/jwt"
	"github.com/abisatria1/simple-dating-app/src/domain/user"
	"github.com/abisatria1/simple-dating-app/src/model"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type AuthUseCaseManager interface {
	Login(ctx context.Context, payload model.LoginRequest) (string, error)
}

type Options struct {
	UserDomain user.UserDomainManager
	JwtManager jwt.JwtManager
}

type authUseCase struct {
	UserDomain user.UserDomainManager
	JwtManager jwt.JwtManager
}

func New(o *Options) AuthUseCaseManager {
	return &authUseCase{
		UserDomain: o.UserDomain,
		JwtManager: o.JwtManager,
	}
}

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
	if jwtToken, err = m.JwtManager.GenerateUserJwt(user.ID); err != nil {
		err = errors.Wrapf(err, "error generate jwt: %s", err.Error())
		return
	}
	return
}
