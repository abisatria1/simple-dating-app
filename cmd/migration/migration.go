package main

import (
	"log"

	"github.com/abisatria1/simple-dating-app/pkg/gorm"
	"github.com/abisatria1/simple-dating-app/src/config"
	"github.com/abisatria1/simple-dating-app/src/domain/migration"
)

func main() {
	cfg := &config.MainConfig{
		App: config.App{
			Port:                3999,
			DefaultWriteTimeout: 10,
			DefaultReadTimeout:  10,
		},
		DB: gorm.DbConfig{
			Dsn:           "root:password@tcp(127.0.0.1:3306)/simple_bumble?charset=utf8mb4&parseTime=True&loc=Local",
			RetryInterval: 10,
			MaxIdleCon:    10,
			MaxCon:        100,
		},
	}
	db := gorm.New(cfg.DB)
	migrationManager := migration.NewGormMigration(&migration.Options{
		DB: db,
	})
	err := migrationManager.DoMigration()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successfuly do migration")
}
