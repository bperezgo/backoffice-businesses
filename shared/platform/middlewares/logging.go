package middlewares

import (
	"github.com/bperezgo/backoffice-businesses/shared/platform/handler"
	"github.com/bperezgo/backoffice-businesses/shared/platform/handlertypes"
	"github.com/bperezgo/backoffice-businesses/shared/platform/logger"
)

type LoggerMiddleware struct {
	logger  *logger.Logger
	handler handler.Handler
}

func NewLoggerMiddleware(handler handler.Handler) *LoggerMiddleware {
	return &LoggerMiddleware{
		logger:  logger.GetLogger(),
		handler: handler,
	}
}

func (h *LoggerMiddleware) GetMethod() handler.HandlerMethod {
	return h.handler.GetMethod()
}

func (h *LoggerMiddleware) GetPath() string {
	return h.handler.GetPath()
}

func (h *LoggerMiddleware) Function(req handlertypes.Request) handlertypes.Response {
	h.logger.Info(logger.LogInput{
		Action: "REQUEST",
		State:  logger.SUCCESS,
		Http: &logger.LogHttpInput{
			Request: req,
		},
		Meta: req.Meta,
	})

	res := h.handler.Function(req)

	h.logger.Info(logger.LogInput{
		Action: "RESPONSE",
		State:  logger.SUCCESS,
		Http: &logger.LogHttpInput{
			Request:  req,
			Response: res,
		},
		Meta: req.Meta,
	})

	return res
}

func (h *LoggerMiddleware) GetEmptyRequest() handlertypes.Request {
	return h.handler.GetEmptyRequest()
}
