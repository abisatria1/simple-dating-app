package gorm

import "gorm.io/gorm"

type DbConfig struct {
	Dsn           string `env:"dsn"`
	RetryInterval int    `env:"retry_interval"`
	MaxIdleCon    int    `env:"max_idle_con"`
	MaxCon        int    `env:"max_con"`
}

type Transaction interface {
	Commit() *gorm.DB
	Rollback() *gorm.DB
}
