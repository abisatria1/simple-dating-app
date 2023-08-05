package feeds

import (
	"context"

	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/domain/user"
	"github.com/abisatria1/simple-dating-app/src/model"
)

type FeedsUsecaseManager interface {
	GetFeedsForUser(ctx context.Context, user entity.User, param model.GetFeedsForUserParam) (userFeeds []model.UserFeed, err error)
	Dislike(ctx context.Context, currentUser entity.User, targetUserID int64) (err error)
	Like(ctx context.Context, currentUser entity.User, targetUserID int64) (err error)
}

type Options struct {
	UserDomain user.UserDomainManager
}

type feedsUseCase struct {
	UserDomain user.UserDomainManager
}

func New(o *Options) FeedsUsecaseManager {
	return &feedsUseCase{
		UserDomain: o.UserDomain,
	}
}
