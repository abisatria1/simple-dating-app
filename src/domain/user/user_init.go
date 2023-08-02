package user

import (
	"context"

	"github.com/abisatria1/simple-dating-app/src/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserDomainManager interface {
	GetUserByEmail(ctx context.Context, email string) (user *model.User, err error)
	CreateUser(ctx context.Context, user model.User, tx *gorm.DB) (ID int64, err error)
}

type UserDomain struct {
	DB *gorm.DB
}

type Options struct {
	DB *gorm.DB
}

func New(o *Options) UserDomainManager {
	return &UserDomain{
		DB: o.DB,
	}
}

func (u *UserDomain) GetUserByEmail(ctx context.Context, email string) (user *model.User, err error) {
	user = &model.User{}
	if err = u.DB.WithContext(ctx).Where(&model.User{Email: email}).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	return
}

func (u *UserDomain) CreateUser(ctx context.Context, user model.User, tx *gorm.DB) (ID int64, err error) {
	executor := u.DB
	if tx != nil {
		executor = tx
	}

	if err = executor.WithContext(ctx).Create(&user).Error; err != nil {
		return
	}
	ID = user.ID
	return
}
