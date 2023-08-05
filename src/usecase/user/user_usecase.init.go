package user

import (
	"context"

	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/domain/interest"
	"github.com/abisatria1/simple-dating-app/src/domain/user"
	"github.com/abisatria1/simple-dating-app/src/model"
)

type UserUseCaseManager interface {
	RegisterUser(ctx context.Context, payload model.RegisterUserRequest) (ID int64, err error)
	InsertUserInterest(ctx context.Context, userID int64, payload model.InsertUserInterestRequest) (affectedRows int64, err error)
	GetUserMatchList(ctx context.Context, user entity.User) (matchUsers []model.UserFeed, err error)
	GetUserPotentialMatchList(ctx context.Context, user entity.User) (potentialUsers []model.UserFeed, err error)
	GetAllInterest(ctx context.Context) (interests []entity.Interest, err error)
}

type Options struct {
	UserDomain     user.UserDomainManager
	InterestDomain interest.InterestDomainManager
}

type userUseCase struct {
	UserDomain     user.UserDomainManager
	InterestDomain interest.InterestDomainManager
}

func New(o *Options) UserUseCaseManager {
	return &userUseCase{
		UserDomain:     o.UserDomain,
		InterestDomain: o.InterestDomain,
	}
}
