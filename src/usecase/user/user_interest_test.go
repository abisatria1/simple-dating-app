package user

import (
	"context"
	"errors"
	"testing"

	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/domain/interest"
	"github.com/abisatria1/simple-dating-app/src/domain/user"
	"github.com/abisatria1/simple-dating-app/src/model"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

type testObject struct {
	userDomain     *user.MockUserDomainManager
	interestDomain *interest.MockInterestDomainManager
	module         *userUseCase
}

func (obj *testObject) MockRunAsExpected(t *testing.T) {
	obj.userDomain.AssertExpectations(t)
	obj.interestDomain.AssertExpectations(t)
}

func doTest(t *testing.T, fn func(*GomegaWithT, *testObject)) {
	mockUserDomain := user.NewMockUserDomainManager(t)
	mockInterestDomain := interest.NewMockInterestDomainManager(t)

	obj := testObject{
		userDomain:     mockUserDomain,
		interestDomain: mockInterestDomain,
	}

	module := userUseCase{
		UserDomain:     mockUserDomain,
		InterestDomain: mockInterestDomain,
	}

	obj.module = &module
	defer obj.MockRunAsExpected(t)

	g := NewGomegaWithT(t)
	fn(g, &obj)
}

func Test_userUseCase_GetAllInterest(t *testing.T) {
	t.Run("error happen", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			obj.interestDomain.On("GetAllInterests", mock.Anything).Return([]entity.Interest{}, errors.New("error"))
			_, err := obj.module.GetAllInterest(context.TODO())
			g.Expect(err).Should(HaveOccurred())
		})
	})
	t.Run("success", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			obj.interestDomain.On("GetAllInterests", mock.Anything).Return([]entity.Interest{}, nil)
			_, err := obj.module.GetAllInterest(context.TODO())
			g.Expect(err).ShouldNot(HaveOccurred())
		})
	})
}

func Test_userUseCase_InsertUserInterest(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			userID := int64(1)
			obj.userDomain.On("BulkInsertUserInterest", mock.Anything, mock.Anything, mock.Anything).Return(int64(0), errors.New("error"))
			_, err := obj.module.InsertUserInterest(context.Background(), userID, model.InsertUserInterestRequest{})
			g.Expect(err).Should(HaveOccurred())
		})
	})
	t.Run("success", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			userID := int64(1)
			obj.userDomain.On("BulkInsertUserInterest", mock.Anything, []entity.UserInterest{
				{
					UserID:     userID,
					InterestID: 1,
				},
				{
					UserID:     userID,
					InterestID: 2,
				},
				{
					UserID:     userID,
					InterestID: 3,
				},
			}, mock.Anything).Return(int64(1), nil)
			_, err := obj.module.InsertUserInterest(context.Background(), userID, model.InsertUserInterestRequest{
				InterestIds: []int64{1, 2, 3},
			})
			g.Expect(err).ShouldNot(HaveOccurred())
		})
	})
}
