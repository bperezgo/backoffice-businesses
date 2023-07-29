package middlewares

import (
	"github.com/bperezgo/backoffice-businesses/shared/platform/handler"
	"github.com/bperezgo/backoffice-businesses/shared/platform/handlertypes"
	"github.com/google/uuid"
)

type MetadataMiddleware struct {
	handler handler.Handler
}

func NewMetadataMiddleware(handler handler.Handler) *MetadataMiddleware {
	return &MetadataMiddleware{
		handler: handler,
	}
}

func (h *MetadataMiddleware) GetMethod() handler.HandlerMethod {
	return h.handler.GetMethod()
}

func (h *MetadataMiddleware) GetPath() string {
	return h.handler.GetPath()
}

func (h *MetadataMiddleware) Function(req handlertypes.Request) handlertypes.Response {
	req.Meta = &handlertypes.Meta{
		RequestId: uuid.NewString(),
	}

	return h.handler.Function(req)
}

func (h *MetadataMiddleware) GetEmptyRequest() handlertypes.Request {
	return h.handler.GetEmptyRequest()
}
