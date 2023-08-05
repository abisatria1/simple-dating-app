package main

import (
	"log"

	"github.com/abisatria1/simple-dating-app/pkg/gorm"
	"github.com/abisatria1/simple-dating-app/src/config"
	"github.com/abisatria1/simple-dating-app/src/domain/seed"
)

func main() {
	cfg := &config.MainConfig{}
	config.Init(cfg)

	db := gorm.New(cfg.DB)
	migrationManager := seed.NewGormSeeder(&seed.Options{
		DB: db,
	})
	err := migrationManager.DoSeeding()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successfuly do seeding")
}
