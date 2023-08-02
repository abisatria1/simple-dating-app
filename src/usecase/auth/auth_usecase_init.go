package auth

import (
	"context"

	"github.com/abisatria1/simple-dating-app/src/domain/user"
	"github.com/abisatria1/simple-dating-app/src/model"
)

type AuthUseCaseManager interface {
	Login(ctx context.Context, payload model.LoginRequest) (string, error)
	RegisterUser(ctx context.Context, payload model.RegisterUserRequest) (ID int64, err error)
}

type Options struct {
	UserDomain user.UserDomainManager
}

type authUseCase struct {
	UserDomain user.UserDomainManager
}

func New(o *Options) AuthUseCaseManager {
	return &authUseCase{
		UserDomain: o.UserDomain,
	}
}
