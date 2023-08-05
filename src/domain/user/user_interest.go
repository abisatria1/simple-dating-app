package user

import (
	"context"

	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (u *UserDomain) BulkInsertUserInterest(ctx context.Context, userInterest []entity.UserInterest, tx *gorm.DB) (affectedRows int64, err error) {
	executor := u.DB
	if tx != nil {
		executor = tx
	}
	query := executor.WithContext(ctx).Create(&userInterest)
	if err = query.Error; err != nil {
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	affectedRows = query.RowsAffected
	return
}
