package config

import (
	"fmt"
	"log"

	env "github.com/Netflix/go-env"
	"github.com/abisatria1/simple-dating-app/pkg/gorm"
	"github.com/joho/godotenv"
)

type MainConfig struct {
	App App
	DB  gorm.DbConfig
	Jwt JwtConfig
}

func Init(cfg *MainConfig) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	if result, err := env.UnmarshalFromEnviron(cfg); err != nil {
		fmt.Print(result)
		fmt.Print(err)
		log.Fatalf("err init config cause : %s", err.Error())
	}
}

type (
	App struct {
		Port                int `env:"port"`
		DefaultWriteTimeout int `env:"default_write_timeout"`
		DefaultReadTimeout  int `env:"default_read_timeout"`
	}

	JwtConfig struct {
		SignKey    string `env:"sign_key"`
		Expiration int64  `env:"expiration"`
	}
)
