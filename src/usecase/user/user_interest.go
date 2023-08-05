package user

import (
	"context"

	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/model"
	"github.com/pkg/errors"
)

func (m *userUseCase) GetAllInterest(ctx context.Context) (interests []entity.Interest, err error) {
	return m.InterestDomain.GetAllInterests(ctx)
}

func (m *userUseCase) InsertUserInterest(ctx context.Context, userID int64, payload model.InsertUserInterestRequest) (affectedRows int64, err error) {
	userInterests := []entity.UserInterest{}
	for _, interest := range payload.InterestIds {
		m := entity.UserInterest{
			UserID:     userID,
			InterestID: interest,
		}
		userInterests = append(userInterests, m)
	}

	if affectedRows, err = m.UserDomain.BulkInsertUserInterest(ctx, userInterests, nil); err != nil {
		err = errors.Wrapf(err, "[InsertUserInterest] BulkInsertUserInterest error : %s", err.Error())
		return
	}

	return
}
