package services

import (
	"context"

	"github.com/bperezgo/backoffice-businesses/models"
	"github.com/bperezgo/backoffice-businesses/ports"
)

type UserService struct {
	userRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Create(ctx context.Context) error {
	userModel := models.User{}
	return s.userRepository.Save(ctx, userModel)
}
