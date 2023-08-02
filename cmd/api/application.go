package main

import (
	"fmt"

	"github.com/abisatria1/simple-dating-app/pkg/echohttp"
	"github.com/abisatria1/simple-dating-app/pkg/gorm"
	"github.com/abisatria1/simple-dating-app/src/config"
	"github.com/abisatria1/simple-dating-app/src/handler"
	"github.com/abisatria1/simple-dating-app/src/service"
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
	httpServer := echohttp.New(&echohttp.Options{
		ListenAddress: cfg.App.Port,
		WriteTimeout:  cfg.App.DefaultWriteTimeout,
		ReadTimeout:   cfg.App.DefaultReadTimeout,
	})
	service := service.New(&service.Options{
		Config: cfg,
	})
	handlers := handler.New(service)
	handlers.RegisterHttpHandlers(httpServer.GetEcho())

	go httpServer.Run()
	err := <-httpServer.ListenError()
	if err != nil {
		fmt.Printf("Error starting web server, exiting gracefully: %v", err)
	}
}
