package gorm

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(cfg DbConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("can't connect to database, cause : %s", err.Error())
	}
	return db
}
