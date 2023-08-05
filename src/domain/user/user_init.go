package user

import (
	"context"
	"time"

	"github.com/abisatria1/simple-dating-app/src/constant"
	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/model"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserDomainManager interface {
	GetUserByID(ctx context.Context, ID int64) (user *entity.User, err error)
	GetUserByEmail(ctx context.Context, email string) (user *entity.User, err error)
	GetUserFeeds(ctx context.Context, userID int64) (users []model.UserFeed, err error)
	CreateUser(ctx context.Context, user entity.User, tx *gorm.DB) (ID int64, err error)
	BulkInsertUserInterest(ctx context.Context, userInterest []entity.UserInterest, tx *gorm.DB) (affectedRows int64, err error)
	ValidateUserQuota(ctx context.Context, user entity.User) (valid bool)
	IsLikedBy(ctx context.Context, userID int64, likeBy int64) (userLike *entity.UserLike, err error)
	InsertLike(ctx context.Context, like entity.UserLike, tx *gorm.DB) (likeID int64, err error)
	DeleteLike(ctx context.Context, criteria entity.UserLike, tx *gorm.DB) (resultAffected int64, err error)
	UpdateUser(ctx context.Context, userID int64, updateField entity.User, tx *gorm.DB) (resultAffected int64, err error)
	UpdateQuota(ctx context.Context, userID, quota int64, tx *gorm.DB) (resultAffected int64, err error)
	InsertBlacklistUser(ctx context.Context, blacklist entity.UserBlacklist, tx *gorm.DB) (blacklistID int64, err error)
	GetIsUserBlacklistedFromCurrentUserFeeds(ctx context.Context, currentUser, targetUser int64) (userBlacklist *entity.UserBlacklist, err error)
	InsertMatch(ctx context.Context, match entity.UserMatch, tx *gorm.DB) (matchID int64, err error)
	GetMatchUsers(ctx context.Context, userID int64) (users []model.UserFeed, err error)
	GetPotentialMatchUsers(ctx context.Context, userID int64) (users []model.UserFeed, err error)
	Begin(ctx context.Context) (tx *gorm.DB)
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

func (u *UserDomain) GetUserByID(ctx context.Context, ID int64) (user *entity.User, err error) {
	user = &entity.User{}
	if err = u.DB.WithContext(ctx).Where(&entity.User{ID: ID}).Preload("Interests").First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	return
}

func (u *UserDomain) GetUserByEmail(ctx context.Context, email string) (user *entity.User, err error) {
	user = &entity.User{}
	if err = u.DB.WithContext(ctx).Where(&entity.User{Email: email}).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	return
}

func (u *UserDomain) CreateUser(ctx context.Context, user entity.User, tx *gorm.DB) (ID int64, err error) {
	executor := u.DB
	if tx != nil {
		executor = tx
	}
	if err = executor.WithContext(ctx).Create(&user).Error; err != nil {
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	ID = user.ID
	return
}

func (u *UserDomain) GetUserFeeds(ctx context.Context, userID int64) (users []model.UserFeed, err error) {
	allUsers := []entity.User{}
	query := u.DB.WithContext(ctx)
	query = query.
		Model(&entity.User{}).
		Joins("LEFT JOIN user_blacklists on users.id = user_blacklists.target_user_id and user_blacklists.user_id = ? and user_blacklists.expired_at >= ?", userID, time.Now()).
		Where("users.id <> ? and user_blacklists.target_user_id is null", userID).
		Order("RAND()").
		Preload("Interests").
		Find(&allUsers)
	if err = query.Error; err != nil {
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}

	copier.Copy(&users, &allUsers)

	return
}

func (u *UserDomain) ValidateUserQuota(ctx context.Context, user entity.User) (valid bool) {
	if user.UserType == constant.UserTypePremium {
		valid = true
		return
	}
	valid = user.Quota > 0
	return
}

func (u *UserDomain) Begin(ctx context.Context) (tx *gorm.DB) {
	return u.DB.Begin()
}

func (u *UserDomain) IsLikedBy(ctx context.Context, userID int64, likeBy int64) (userLike *entity.UserLike, err error) {
	userLike = &entity.UserLike{}
	if err = u.DB.WithContext(ctx).Where("`like` = ? and like_by = ? and expired_at > ?", userID, likeBy, time.Now()).First(userLike).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	return
}

func (u *UserDomain) InsertLike(ctx context.Context, like entity.UserLike, tx *gorm.DB) (likeID int64, err error) {
	executor := u.DB
	if tx != nil {
		executor = tx
	}
	query := executor.WithContext(ctx).Create(&like)
	if err = query.Error; err != nil {
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	likeID = like.ID
	return
}

func (u *UserDomain) DeleteLike(ctx context.Context, criteria entity.UserLike, tx *gorm.DB) (resultAffected int64, err error) {
	executor := u.DB
	if tx != nil {
		executor = tx
	}
	query := executor.WithContext(ctx).Delete(&criteria)
	if err = query.Error; err != nil {
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	resultAffected = query.RowsAffected
	return
}

func (u *UserDomain) UpdateUser(ctx context.Context, userID int64, updateField entity.User, tx *gorm.DB) (resultAffected int64, err error) {
	executor := u.DB
	if tx != nil {
		executor = tx
	}
	query := executor.WithContext(ctx).Model(&entity.User{}).Where("id = ?", userID).Updates(&updateField)
	if err = query.Error; err != nil {
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	resultAffected = query.RowsAffected
	return
}

func (u *UserDomain) UpdateQuota(ctx context.Context, userID, quota int64, tx *gorm.DB) (resultAffected int64, err error) {
	executor := u.DB
	if tx != nil {
		executor = tx
	}
	query := executor.WithContext(ctx).Model(&entity.User{}).Select("quota").Where("id = ?", userID).Update("quota", quota)
	if err = query.Error; err != nil {
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	resultAffected = query.RowsAffected
	return
}

func (u *UserDomain) InsertBlacklistUser(ctx context.Context, blacklist entity.UserBlacklist, tx *gorm.DB) (blacklistID int64, err error) {
	executor := u.DB
	if tx != nil {
		executor = tx
	}
	query := executor.WithContext(ctx).Create(&blacklist)
	if err = query.Error; err != nil {
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	blacklistID = blacklist.ID
	return
}

func (u *UserDomain) GetIsUserBlacklistedFromCurrentUserFeeds(ctx context.Context, currentUser, targetUser int64) (userBlacklist *entity.UserBlacklist, err error) {
	userBlacklist = &entity.UserBlacklist{}
	if err = u.DB.WithContext(ctx).Where("user_id = ? and target_user_id = ? and expired_at > ?", currentUser, targetUser, time.Now()).First(userBlacklist).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	return
}

func (u *UserDomain) InsertMatch(ctx context.Context, match entity.UserMatch, tx *gorm.DB) (matchID int64, err error) {
	executor := u.DB
	if tx != nil {
		executor = tx
	}
	query := executor.WithContext(ctx).Create(&match)
	if err = query.Error; err != nil {
		err = errors.Wrapf(err, "query execution error : %s", err.Error())
		return
	}
	matchID = match.ID
	return
}
