package user

import (
	"context"
	"time"

	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/model"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
)

func (u *UserDomain) GetMatchUsers(ctx context.Context, userID int64) (users []model.UserFeed, err error) {
	allUsers := []entity.User{}
	users = []model.UserFeed{}
	query := u.DB.WithContext(ctx)
	query = query.
		Model(&entity.User{}).
		Joins("LEFT JOIN user_matches on users.id = user_matches.target_user_id").
		Where("user_matches.user_id = ? and user_matches.expired_at >= ?", userID, time.Now()).
		Preload("Interests").
		Find(&allUsers)
	if err = query.Error; err != nil {
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}

	copier.Copy(&users, &allUsers)
	return
}

func (u *UserDomain) GetPotentialMatchUsers(ctx context.Context, userID int64) (users []model.UserFeed, err error) {
	allUsers := []entity.User{}
	users = []model.UserFeed{}
	query := u.DB.WithContext(ctx)
	query = query.
		Model(&entity.User{}).
		Joins("LEFT JOIN user_likes on users.id = user_likes.like_by").
		Where("user_likes.`like` = ? and user_likes.expired_at >= ?", userID, time.Now()).
		Preload("Interests").
		Find(&allUsers)
	if err = query.Error; err != nil {
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	copier.Copy(&users, &allUsers)
	return
}
