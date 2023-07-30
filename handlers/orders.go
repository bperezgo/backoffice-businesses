package handlers

import (
	"context"
	"net/http"

	"github.com/bperezgo/backoffice-businesses/shared/platform/handler"
	"github.com/bperezgo/backoffice-businesses/shared/platform/handlertypes"
)

type CreateOrderResponse struct {
	Message string `json:"message"`
}

type CreateOrderRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateOrderHandler struct {
	handler.BasisHandler[CreateOrderRequest]
}

func NewCreateOrderHandler() *CreateOrderHandler {
	return &CreateOrderHandler{
		BasisHandler: handler.BasisHandler[CreateOrderRequest]{
			HandlerMethod: handler.POST,
			Path:          "/orders",
			BasisBody:     CreateOrderRequest{},
		},
	}
}

func (h *CreateOrderHandler) Function(ctx context.Context, req handlertypes.Request) handlertypes.Response {
	return handlertypes.Response{
		Body: CreateOrderResponse{
			Message: "Order created",
		},
		HttpStatus: http.StatusCreated,
	}
}
