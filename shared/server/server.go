package server

import (
	"fmt"
	"net/http"

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

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
		err, ok := recovered.(string)
		if ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

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
