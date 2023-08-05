package user

import (
	"context"
	"net/http"

	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/model"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (m *userUseCase) RegisterUser(ctx context.Context, payload model.RegisterUserRequest) (ID int64, err error) {
	user, err := payload.ToModel()
	if err != nil {
		err = errors.Wrapf(err, "[RegisterUser] convert payload err: %s", err.Error())
		return
	}

	if err = user.HashPassword(); err != nil {
		err = errors.Wrapf(err, "[RegisterUser] hashed password err: %s", err.Error())
		return
	}

	checkEmail, err := m.UserDomain.GetUserByEmail(ctx, user.Email)
	if err != nil {
		err = errors.Wrapf(err, "[RegisterUser] GetUserByEmail error : %s", err.Error())
		return
	}
	if checkEmail != nil {
		err = echo.NewHTTPError(http.StatusBadRequest, "email already used")
		return
	}

	if ID, err = m.UserDomain.CreateUser(ctx, user, nil); err != nil {
		err = errors.Wrapf(err, "[RegisterUser] CreateUser error : %s", err.Error())
		return
	}

	return
}

func (m *userUseCase) GetUserMatchList(ctx context.Context, user entity.User) (matchUsers []model.UserFeed, err error) {
	matchUsers, err = m.UserDomain.GetMatchUsers(ctx, user.ID)
	if err != nil {
		err = errors.Wrapf(err, "[GetUserMatchList] GetMatchUsers error : %s", err.Error())
		return
	}
	return
}

func (m *userUseCase) GetUserPotentialMatchList(ctx context.Context, user entity.User) (potentialUsers []model.UserFeed, err error) {
	potentialUsers, err = m.UserDomain.GetPotentialMatchUsers(ctx, user.ID)
	if err != nil {
		err = errors.Wrapf(err, "[GetUserPotentialMatchList] GetPotentialMatchUsers error : %s", err.Error())
		return
	}
	return
}
