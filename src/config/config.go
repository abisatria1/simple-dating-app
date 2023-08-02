package config

import "github.com/abisatria1/simple-dating-app/pkg/gorm"

type MainConfig struct {
	App App
	DB  gorm.DbConfig
}

type (
	App struct {
		Port                int `json:"port"`
		DefaultWriteTimeout int `json:"default_write_timeout"`
		DefaultReadTimeout  int `json:"default_read_timeout"`
	}
)
