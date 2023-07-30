package handlers

import (
	"context"
	"net/http"

	"github.com/bperezgo/backoffice-businesses/shared/platform/handler"
	"github.com/bperezgo/backoffice-businesses/shared/platform/handlertypes"
)

type GetHealthResponse struct {
	Message string `json:"message"`
	handler.BasisHandler[interface{}]
}

type GetHealthHandler struct {
	handler.BasisHandler[interface{}]
}

func NewGetHealthHandler() *GetHealthHandler {
	return &GetHealthHandler{
		BasisHandler: handler.BasisHandler[interface{}]{
			HandlerMethod: handler.GET,
			Path:          "/health",
			BasisBody:     nil,
		},
	}
}

func (h *GetHealthHandler) Function(ctx context.Context, req handlertypes.Request) handlertypes.Response {
	return handlertypes.Response{
		Body: GetHealthResponse{
			Message: "Server is up and running",
		},
		HttpStatus: http.StatusOK,
	}
}
