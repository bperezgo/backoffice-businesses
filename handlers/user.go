package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bperezgo/backoffice-businesses/services"
	"github.com/bperezgo/backoffice-businesses/shared/platform/handler"
	"github.com/bperezgo/backoffice-businesses/shared/platform/handlertypes"
)

type CreateUserResponse struct {
	Message string `json:"message"`
}

type CreateUserRequest struct {
	Email            string `json:"email"`
	Name             string `json:"name"`
	LastName         string `json:"lastName"`
	Phone            string `json:"phone"`
	CountryPhoneCode string `json:"countryPhoneCode"`
}

type CreateUserHandler struct {
	handler.BasisHandler[CreateUserRequest]
	userService *services.UserService
}

func NewCreateUserHandler(userService *services.UserService) *CreateUserHandler {
	return &CreateUserHandler{
		BasisHandler: handler.BasisHandler[CreateUserRequest]{
			HandlerMethod: handler.POST,
			Path:          "/users",
			BasisBody:     CreateUserRequest{},
		},
	}
}

func (h *CreateUserHandler) Function(ctx context.Context, req handlertypes.Request) handlertypes.Response {
	userBody, ok := req.Body.(CreateUserRequest)
	if !ok {
		return handlertypes.Response{
			Body:       nil,
			HttpStatus: http.StatusBadRequest,
		}
	}
	fmt.Println(userBody.Email)
	err := h.userService.Create(ctx)
	if err != nil {
		return handlertypes.Response{
			Body:       nil,
			HttpStatus: http.StatusBadRequest,
		}
	}

	return handlertypes.Response{
		Body: CreateOrderResponse{
			Message: "User created",
		},
		HttpStatus: http.StatusCreated,
	}
}
