package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/abisatria1/simple-dating-app/pkg/jwt"
	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"github.com/abisatria1/simple-dating-app/src/domain/user"
	"github.com/abisatria1/simple-dating-app/src/model"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

func Test_authUseCase_Login(t *testing.T) {
	t.Run("Error get user email", func(t *testing.T) {
		g := NewGomegaWithT(t)
		email := "testemail@gmail.com"
		mockUserDomain := user.NewMockUserDomainManager(t)
		mockUserDomain.On("GetUserByEmail", mock.Anything, email).Return(nil, errors.New("error"))

		mo := authUseCase{
			UserDomain: mockUserDomain,
		}

		_, err := mo.Login(context.Background(), model.LoginRequest{Email: email, Password: "password"})
		g.Expect(err).Should(HaveOccurred())
	})
	t.Run("Invalid email", func(t *testing.T) {
		g := NewGomegaWithT(t)
		email := "invalid@gmail.com"
		mockUserDomain := user.NewMockUserDomainManager(t)
		mockUserDomain.On("GetUserByEmail", mock.Anything, email).Return(nil, nil)

		mo := authUseCase{
			UserDomain: mockUserDomain,
		}

		_, err := mo.Login(context.Background(), model.LoginRequest{Email: email, Password: "password"})
		g.Expect(err).Should(HaveOccurred())
	})
	t.Run("Invalid password", func(t *testing.T) {
		g := NewGomegaWithT(t)
		email := "test@gmail.com"
		mockUserDomain := user.NewMockUserDomainManager(t)
		user := &entity.User{Password: "realpassword"}
		user.HashPassword()

		mockUserDomain.On("GetUserByEmail", mock.Anything, email).Return(user, nil)

		mo := authUseCase{
			UserDomain: mockUserDomain,
		}

		_, err := mo.Login(context.Background(), model.LoginRequest{Email: email, Password: "password"})
		g.Expect(err).Should(HaveOccurred())
		mockUserDomain.AssertExpectations(t)
	})
	t.Run("Success", func(t *testing.T) {
		g := NewGomegaWithT(t)
		email := "test@gmail.com"
		password := "password"
		mockUserDomain := user.NewMockUserDomainManager(t)
		mockJwtManager := jwt.NewMockJwtManager(t)
		user := &entity.User{ID: 1, Password: password, Email: email}
		user.HashPassword()

		mockUserDomain.On("GetUserByEmail", mock.Anything, email).Return(user, nil)
		mockJwtManager.On("GenerateUserJwt", user.ID).Return("token", nil)

		mo := authUseCase{
			UserDomain: mockUserDomain,
			JwtManager: mockJwtManager,
		}

		token, err := mo.Login(context.Background(), model.LoginRequest{Email: email, Password: password})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(token).Should(Equal("token"))
		mockUserDomain.AssertExpectations(t)
		mockJwtManager.AssertExpectations(t)
	})

}
