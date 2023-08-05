package user

import (
	"context"
	"errors"
	"testing"

	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/model"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

func Test_userUseCase_RegisterUser(t *testing.T) {
	t.Run("error GetUserByEmail", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			user := model.RegisterUserRequest{
				Email: "testingemail@gmail.com",
			}
			obj.userDomain.On("GetUserByEmail", mock.Anything, user.Email).Return(nil, errors.New("error"))
			_, err := obj.module.RegisterUser(context.Background(), user)
			g.Expect(err).Should(HaveOccurred())
		})
	})
	t.Run("duplicate email", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			user := model.RegisterUserRequest{
				Email: "testingemail@gmail.com",
			}
			obj.userDomain.On("GetUserByEmail", mock.Anything, user.Email).Return(&entity.User{}, nil)
			_, err := obj.module.RegisterUser(context.Background(), user)
			g.Expect(err).Should(HaveOccurred())
		})
	})
	t.Run("error create user", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			user := model.RegisterUserRequest{
				Email: "testingemail@gmail.com",
			}
			obj.userDomain.On("GetUserByEmail", mock.Anything, user.Email).Return(nil, nil)
			obj.userDomain.On("CreateUser", mock.Anything, mock.Anything, mock.Anything).Return(int64(0), errors.New("error"))
			_, err := obj.module.RegisterUser(context.Background(), user)
			g.Expect(err).Should(HaveOccurred())
		})
	})
	t.Run("Success", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			user := model.RegisterUserRequest{
				Email:    "testingemail@gmail.com",
				Password: "password",
			}
			obj.userDomain.On("GetUserByEmail", mock.Anything, user.Email).Return(nil, nil)
			obj.userDomain.On("CreateUser", mock.Anything, mock.Anything, mock.Anything).Return(int64(12), nil)
			ID, err := obj.module.RegisterUser(context.Background(), user)
			g.Expect(err).ShouldNot(HaveOccurred())
			g.Expect(ID).Should(Equal(int64(12)))
		})
	})
}

func Test_userUseCase_GetUserMatchList(t *testing.T) {
	t.Run("Error happen", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			obj.userDomain.On("GetMatchUsers", mock.Anything, mock.Anything).Return([]model.UserFeed{}, errors.New("error"))
			_, err := obj.module.GetUserMatchList(context.Background(), entity.User{})
			g.Expect(err).Should(HaveOccurred())
		})
	})
	t.Run("Success", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			obj.userDomain.On("GetMatchUsers", mock.Anything, mock.Anything).Return([]model.UserFeed{}, nil)
			result, err := obj.module.GetUserMatchList(context.Background(), entity.User{})
			g.Expect(err).ShouldNot(HaveOccurred())
			g.Expect(result).Should(Equal([]model.UserFeed{}))
		})
	})
}

func Test_userUseCase_GetUserPotentialMatchList(t *testing.T) {
	t.Run("Error happen", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			obj.userDomain.On("GetPotentialMatchUsers", mock.Anything, mock.Anything).Return([]model.UserFeed{}, errors.New("error"))
			_, err := obj.module.GetUserPotentialMatchList(context.Background(), entity.User{})
			g.Expect(err).Should(HaveOccurred())
		})
	})
	t.Run("Success", func(t *testing.T) {
		doTest(t, func(g *GomegaWithT, obj *testObject) {
			obj.userDomain.On("GetPotentialMatchUsers", mock.Anything, mock.Anything).Return([]model.UserFeed{}, nil)
			result, err := obj.module.GetUserPotentialMatchList(context.Background(), entity.User{})
			g.Expect(err).ShouldNot(HaveOccurred())
			g.Expect(result).Should(Equal([]model.UserFeed{}))
		})
	})
}
