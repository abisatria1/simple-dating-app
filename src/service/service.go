package service

import (
	"github.com/abisatria1/simple-dating-app/pkg/gorm"
	"github.com/abisatria1/simple-dating-app/pkg/jwt"
	"github.com/abisatria1/simple-dating-app/src/config"
	interestDomain "github.com/abisatria1/simple-dating-app/src/domain/interest"
	subscriptionDomain "github.com/abisatria1/simple-dating-app/src/domain/subscription"
	userDomain "github.com/abisatria1/simple-dating-app/src/domain/user"
	"github.com/abisatria1/simple-dating-app/src/middleware"
	"github.com/abisatria1/simple-dating-app/src/usecase/auth"
	"github.com/abisatria1/simple-dating-app/src/usecase/feeds"
	"github.com/abisatria1/simple-dating-app/src/usecase/subscription"
	"github.com/abisatria1/simple-dating-app/src/usecase/user"
)

type Options struct {
	Config *config.MainConfig
}

type Service struct {
	Usecases   Usecases
	Middleware middleware.Middleware
}

type Usecases struct {
	Auth         auth.AuthUseCaseManager
	User         user.UserUseCaseManager
	Feeds        feeds.FeedsUsecaseManager
	Subscription subscription.SubscriptionUseCaseManager
}

func New(o *Options) *Service {
	db := gorm.New(o.Config.DB)
	jwtManager := jwt.New(&jwt.Options{SignKey: o.Config.Jwt.SignKey, Expiration: o.Config.Jwt.Expiration})
	userDomainManager := userDomain.New(&userDomain.Options{DB: db})
	subscriptionDomainManager := subscriptionDomain.New(&subscriptionDomain.Options{DB: db})
	interestDomainManager := interestDomain.New(&interestDomain.Options{DB: db})
	authUseCaseManager := auth.New(&auth.Options{UserDomain: userDomainManager, JwtManager: jwtManager})
	userUseCaseManager := user.New(&user.Options{UserDomain: userDomainManager, InterestDomain: interestDomainManager})
	feedsUseCaseManager := feeds.New(&feeds.Options{UserDomain: userDomainManager})
	subscriptionUseCaseManager := subscription.New(&subscription.Options{UserDomain: userDomainManager, SubscriptionDomain: subscriptionDomainManager})
	middleware := middleware.New(&middleware.Options{UserDomain: userDomainManager, JwtManager: jwtManager, SubscriptionDomain: subscriptionDomainManager})

	return &Service{
		Usecases: Usecases{
			Auth:         authUseCaseManager,
			User:         userUseCaseManager,
			Feeds:        feedsUseCaseManager,
			Subscription: subscriptionUseCaseManager,
		},
		Middleware: middleware,
	}
}
