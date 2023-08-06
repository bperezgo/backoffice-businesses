package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/bperezgo/backoffice-businesses/models"
	"github.com/bperezgo/backoffice-businesses/services"
	"github.com/bperezgo/backoffice-businesses/shared/platform/handler"
	"github.com/bperezgo/backoffice-businesses/shared/platform/handlertypes"
	"github.com/google/uuid"
)

type CreateUserResponse struct {
	Message string `json:"message"`
}

type CreateUserRequest struct {
	Email            string `json:"email" valid:"email,required"`
	Name             string `json:"name" valid:"matches(^[A-Za-zÀ-ÖØ-öø-ÿ\\s]+$),required"`
	LastName         string `json:"lastName" valid:"matches(^[A-Za-zÀ-ÖØ-öø-ÿ\\s]+$),required"`
	Phone            string `json:"phone" valid:"numeric,required"`
	CountryPhoneCode string `json:"countryPhoneCode" valid:"matches(^\\+\\d+$),required"`
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
	user := CreateUserRequest{}
	err := json.Unmarshal(req.Body, &user)

	if err != nil {
		return handlertypes.Response{
			Body:       nil,
			HttpStatus: http.StatusBadRequest,
		}
	}

	err = handler.Validator(&user)
	if err != nil {
		return handlertypes.Response{
			Body:       nil,
			HttpStatus: http.StatusBadRequest,
		}
	}

	nowTime := time.Now()

	userModel := models.User{
		ID:               uuid.NewString(),
		Email:            user.Email,
		Phone:            user.Phone,
		CountryPhoneCode: user.CountryPhoneCode,
		Name:             user.Name,
		LastName:         user.LastName,
		CreatedAt:        nowTime,
		UpdatedAt:        nowTime,
	}

	err = h.userService.Create(ctx, userModel)
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
