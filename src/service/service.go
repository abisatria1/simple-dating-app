package service

import (
	"github.com/abisatria1/simple-dating-app/pkg/gorm"
	"github.com/abisatria1/simple-dating-app/src/config"
	"github.com/abisatria1/simple-dating-app/src/domain/user"
	"github.com/abisatria1/simple-dating-app/src/usecase/auth"
)

type Options struct {
	Config *config.MainConfig
}

type Service struct {
	Usecases Usecases
}

type Usecases struct {
	Auth auth.AuthUseCaseManager
}

func New(o *Options) *Service {
	db := gorm.New(o.Config.DB)
	userDomain := user.New(&user.Options{DB: db})
	authManager := auth.New(&auth.Options{UserDomain: userDomain})

	return &Service{
		Usecases: Usecases{
			Auth: authManager,
		},
	}
}
