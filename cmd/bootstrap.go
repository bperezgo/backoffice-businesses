package main

import (
	"log"

	"github.com/bperezgo/backoffice-businesses/config"
	"github.com/bperezgo/backoffice-businesses/handlers"
	"github.com/bperezgo/backoffice-businesses/shared/platform/logger"
	"github.com/bperezgo/backoffice-businesses/shared/server"

	_ "github.com/lib/pq"
)

func run() error {
	err := config.InitConfig()
	if err != nil {
		log.Fatal("error loading .env file", err)
	}

	c := config.GetConfig()
	logger.InitLogger()
	allHandlers := handlers.InitHandlers(c)
	srv := server.NewServer(allHandlers...)
	return srv.Start(8080)
}
