package ports

import (
	"context"

	"github.com/bperezgo/backoffice-businesses/models"
)

type UserRepository interface {
	Save(ctx context.Context, user models.User) error
}
