package feeds

import (
	"context"
	"time"

	"github.com/abisatria1/simple-dating-app/pkg/custerr"
	"github.com/abisatria1/simple-dating-app/src/constant"
	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/model"
	"github.com/pkg/errors"
)

func (m *feedsUseCase) GetFeedsForUser(ctx context.Context, user entity.User, param model.GetFeedsForUserParam) (userFeeds []model.UserFeed, err error) {
	if userFeeds, err = m.UserDomain.GetUserFeeds(ctx, user.ID); err != nil {
		err = errors.Wrapf(err, "GetUserFeeds error: %s", err.Error())
		return
	}
	return
}

func (m *feedsUseCase) Dislike(ctx context.Context, currentUser entity.User, targetUserID int64) (err error) {
	if currentUser.ID == targetUserID {
		err = custerr.ErrSameAccount
		return
	}

	// check if user is blacklist from current user feeds
	isBlacklist, err := m.UserDomain.GetIsUserBlacklistedFromCurrentUserFeeds(ctx, currentUser.ID, targetUserID)
	if err != nil {
		err = errors.Wrapf(err, "[Dislike] GetIsUserBlacklistedFromCurrentUserFeeds error: %s", err)
		return
	}
	if isBlacklist != nil {
		err = custerr.ErrUserAlreadyAppearInFeeds
		return
	}

	tx := m.UserDomain.Begin(ctx)
	defer tx.Rollback()

	// cek apakah ada yang di dislike melakukan like
	userLike, err := m.UserDomain.IsLikedBy(ctx, currentUser.ID, targetUserID)
	if err != nil {
		err = errors.Wrapf(err, "[Dislike] IsLikedBy error: %s", err)
		return
	}
	// jika melakukan like maka hapus like dari user tersebut
	if userLike != nil {
		criteria := entity.UserLike{
			ID: userLike.ID,
		}
		if _, err = m.UserDomain.DeleteLike(ctx, criteria, tx); err != nil {
			err = errors.Wrapf(err, "[Dislike] DeleteLike error: %s", err)
			return
		}
	}

	if currentUser.UserType == constant.UserTypeNormal {
		if _, err = m.UserDomain.UpdateQuota(ctx, currentUser.ID, currentUser.Quota-1, tx); err != nil {
			err = errors.Wrapf(err, "[Like] UpdateUser error: %s", err)
			return
		}
	}

	// masukan user ke blacklist dan tambahkan 1 hari sebagai expired
	blacklist := entity.UserBlacklist{
		UserID:       currentUser.ID,
		TargetUserID: targetUserID,
		ExpiredAt:    time.Now().Add(constant.DefaultBlacklistExpired * time.Hour),
	}
	if _, err = m.UserDomain.InsertBlacklistUser(ctx, blacklist, tx); err != nil {
		err = errors.Wrapf(err, "[Dislike] InsertBlacklistUser error: %s", err)
		return
	}

	tx.Commit()

	return
}

func (m *feedsUseCase) Like(ctx context.Context, currentUser entity.User, targetUserID int64) (err error) {
	if currentUser.ID == targetUserID {
		err = custerr.ErrSameAccount
		return
	}

	// check if user is blacklist from current user feeds
	isBlacklist, err := m.UserDomain.GetIsUserBlacklistedFromCurrentUserFeeds(ctx, currentUser.ID, targetUserID)
	if err != nil {
		err = errors.Wrapf(err, "[Dislike] GetIsUserBlacklistedFromCurrentUserFeeds error: %s", err)
		return
	}
	if isBlacklist != nil {
		err = custerr.ErrUserAlreadyAppearInFeeds
		return
	}

	tx := m.UserDomain.Begin(ctx)
	defer tx.Rollback()

	// cek apakah ada yang di dislike melakukan like
	userLike, err := m.UserDomain.IsLikedBy(ctx, currentUser.ID, targetUserID)
	if err != nil {
		err = errors.Wrapf(err, "[Like] IsLikedBy error: %s", err)
		return
	}
	// jika melakukan like maka tambahkan user ke dalam match
	if userLike != nil {
		expired := time.Now().Add(constant.DefaultLikeExpired * time.Hour)
		currentMatch := entity.UserMatch{
			UserID:       currentUser.ID,
			TargetUserID: targetUserID,
			ExpiredAt:    expired,
		}
		if _, err = m.UserDomain.InsertMatch(ctx, currentMatch, tx); err != nil {
			err = errors.Wrapf(err, "[Like] InsertLike for current user error: %s", err)
			return
		}

		targetUserMatch := entity.UserMatch{
			UserID:       targetUserID,
			TargetUserID: currentUser.ID,
			ExpiredAt:    expired,
		}
		if _, err = m.UserDomain.InsertMatch(ctx, targetUserMatch, tx); err != nil {
			err = errors.Wrapf(err, "[Like] InsertLike for target user error: %s", err)
			return
		}

		criteria := entity.UserLike{ID: userLike.ID}
		if _, err = m.UserDomain.DeleteLike(ctx, criteria, tx); err != nil {
			err = errors.Wrapf(err, "[Like] DeleteLike error: %s", err)
			return
		}
	} else {
		like := entity.UserLike{
			Like:      targetUserID,
			LikeBy:    currentUser.ID,
			ExpiredAt: time.Now().Add(constant.DefaultLikeExpired * time.Hour),
		}
		if _, err = m.UserDomain.InsertLike(ctx, like, tx); err != nil {
			err = errors.Wrapf(err, "[Like] InsertLike error: %s", err)
			return
		}
	}

	// kurangi kuota
	if currentUser.UserType == constant.UserTypeNormal {
		if _, err = m.UserDomain.UpdateQuota(ctx, currentUser.ID, currentUser.Quota-1, tx); err != nil {
			err = errors.Wrapf(err, "[Like] UpdateUser error: %s", err)
			return
		}
	}

	// masukan user ke blacklist dan tambahkan 1 hari sebagai expired
	blacklist := entity.UserBlacklist{
		UserID:       currentUser.ID,
		TargetUserID: targetUserID,
		ExpiredAt:    time.Now().Add(constant.DefaultBlacklistExpired * time.Hour),
	}
	if _, err = m.UserDomain.InsertBlacklistUser(ctx, blacklist, tx); err != nil {
		err = errors.Wrapf(err, "[Like] InsertBlacklistUser error: %s", err)
		return
	}

	tx.Commit()

	return
}
