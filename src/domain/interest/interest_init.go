package interest

import (
	"context"

	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type InterestDomainManager interface {
	Begin(ctx context.Context) (tx *gorm.DB)
	GetAllInterests(ctx context.Context) (interests []entity.Interest, err error)
}

type InterestDomain struct {
	DB *gorm.DB
}

type Options struct {
	DB *gorm.DB
}

func New(o *Options) InterestDomainManager {
	return &InterestDomain{
		DB: o.DB,
	}
}

func (m *InterestDomain) Begin(ctx context.Context) (tx *gorm.DB) {
	tx = m.DB.Begin()
	return
}

func (m *InterestDomain) GetAllInterests(ctx context.Context) (interests []entity.Interest, err error) {
	interests = []entity.Interest{}
	if err = m.DB.Find(&interests).Error; err != nil {
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	return
}
