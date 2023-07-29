package main

import (
	"github.com/bperezgo/backoffice-businesses/handlers"
	"github.com/bperezgo/backoffice-businesses/shared/platform/logger"
	"github.com/bperezgo/backoffice-businesses/shared/server"
)

func run() error {
	logger.InitLogger()
	allHandlers := handlers.GetHandlers()
	srv := server.NewServer(allHandlers...)
	return srv.Start(8080)
}
