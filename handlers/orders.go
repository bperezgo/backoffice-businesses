package handlers

import (
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
	handler.BasisHandler
}

func NewCreateOrderHandler() *CreateOrderHandler {
	return &CreateOrderHandler{
		BasisHandler: handler.BasisHandler{
			HandlerMethod: handler.POST,
			Path:          "/orders",
			BasisBody:     CreateOrderRequest{},
		},
	}
}

func (h *CreateOrderHandler) Function(req handlertypes.Request) handlertypes.Response {
	return handlertypes.Response{
		Body: CreateOrderResponse{
			Message: "Order created",
		},
		HttpStatus: http.StatusCreated,
	}
}
