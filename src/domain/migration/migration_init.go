package migration

import (
	"github.com/abisatria1/simple-dating-app/src/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MigrationManager interface {
	DoMigration() error
}

type Options struct {
	DB *gorm.DB
}

type GormMigrator struct {
	DB *gorm.DB
}

func NewGormMigration(o *Options) MigrationManager {
	return &GormMigrator{
		DB: o.DB,
	}
}

func (gm *GormMigrator) DoMigration() (err error) {
	err = gm.DB.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		return errors.Wrapf(err, "migrating error : %s", err.Error())
	}
	return
}
