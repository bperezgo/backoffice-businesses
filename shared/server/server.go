package server

import (
	"fmt"

	"github.com/bperezgo/backoffice-businesses/shared/platform/handler"
	"github.com/bperezgo/backoffice-businesses/shared/platform/middlewares"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine

	handlers []handler.Handler
}

func NewServer(handlers ...handler.Handler) *Server {
	r := gin.Default()

	r.Use(gin.Recovery())

	defineGinHandlers(r, handlers)

	return &Server{
		engine:   r,
		handlers: handlers,
	}
}

func (s *Server) Start(port int) error {
	addr := fmt.Sprintf("localhost:%d", port)
	return s.engine.Run(addr)
}

var mapMethods = make(map[handler.HandlerMethod]func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes)

func defineGinHandlers(engine *gin.Engine, handlers []handler.Handler) {
	mapMethods[handler.GET] = engine.GET
	mapMethods[handler.Any] = engine.Any
	mapMethods[handler.DELETE] = engine.DELETE
	mapMethods[handler.HEAD] = engine.HEAD
	mapMethods[handler.POST] = engine.POST
	mapMethods[handler.OPTIONS] = engine.OPTIONS
	mapMethods[handler.PATCH] = engine.PATCH
	mapMethods[handler.PUT] = engine.PUT
	jsonHandler := handler.NewJsonHandler()

	for _, handler := range handlers {
		loggerHandler := middlewares.NewLoggerMiddleware(handler)

		metadataMiddleware := middlewares.NewMetadataMiddleware(loggerHandler)

		ginHandler := jsonHandler.Adapt(metadataMiddleware)

		mapMethods[handler.GetMethod()](handler.GetPath(), ginHandler)
	}
}
