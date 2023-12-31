package main

import (
	"fmt"
	"os"

	"github.com/abisatria1/simple-dating-app/pkg/echohttp"
	"github.com/abisatria1/simple-dating-app/src/config"
	"github.com/abisatria1/simple-dating-app/src/handler"
	"github.com/abisatria1/simple-dating-app/src/service"
)

func main() {
	cfg := &config.MainConfig{}
	config.Init(cfg)
	os.Setenv("TZ", "UTC")

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
