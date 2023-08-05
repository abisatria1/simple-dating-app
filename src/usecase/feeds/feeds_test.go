package feeds

import (
	"context"
	"errors"
	"testing"

	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/domain/user"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type testObject struct {
	userDomain *user.MockUserDomainManager
	module     *feedsUseCase
}

func (obj *testObject) MockRunAsExpected(t *testing.T) {
	obj.userDomain.AssertExpectations(t)
}

func doTest(t *testing.T, fn func(*GomegaWithT, *testObject)) {
	mockUserDomain := user.NewMockUserDomainManager(t)

	obj := testObject{
		userDomain: mockUserDomain,
	}

	module := feedsUseCase{
		UserDomain: mockUserDomain,
	}

	obj.module = &module
	defer obj.MockRunAsExpected(t)

	g := NewGomegaWithT(t)
	fn(g, &obj)
}

type MockGorm gorm.DB

func (g *MockGorm) Rollback() *gorm.DB {
	return nil
}

func (g *MockGorm) Commit() *gorm.DB {
	return nil
}

func Test_feedsUseCase_Dislike(t *testing.T) {
	t.Run("error same user cant dislike each other", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			curr := entity.User{ID: 1}
			target := int64(1)
			err := obj.module.Dislike(context.Background(), curr, target)
			g.Expect(err).Should(HaveOccurred())
		})
	})
	t.Run("error check get user from blacklist", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			curr := entity.User{ID: 1}
			target := int64(2)

			obj.userDomain.On("GetIsUserBlacklistedFromCurrentUserFeeds", mock.Anything, curr.ID, target).Return(nil, errors.New("error"))

			err := obj.module.Dislike(context.Background(), curr, target)
			g.Expect(err).Should(HaveOccurred())
		})
	})
	t.Run("user cant dislike person twice in a day", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			curr := entity.User{ID: 1}
			target := int64(2)

			obj.userDomain.On("GetIsUserBlacklistedFromCurrentUserFeeds", mock.Anything, curr.ID, target).Return(&entity.UserBlacklist{}, nil)

			err := obj.module.Dislike(context.Background(), curr, target)
			g.Expect(err).Should(HaveOccurred())
		})
	})
}
