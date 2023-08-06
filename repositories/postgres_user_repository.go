package repositories

import (
	"context"
	"errors"

	"github.com/bperezgo/backoffice-businesses/models"
	"github.com/bperezgo/backoffice-businesses/shared/platform/logger"
	"github.com/jmoiron/sqlx"
)

type PostgresUserRepository struct {
	db     *sqlx.DB
	logger *logger.Logger
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db:     db,
		logger: logger.GetLogger(),
	}
}

func (r *PostgresUserRepository) Save(ctx context.Context, user models.User) error {
	_, err := r.db.Query(`
	INSERT INTO users (id, country_phone_code, email, name, last_name, phone, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		user.ID,
		user.CountryPhoneCode,
		user.Email,
		user.Name,
		user.LastName,
		user.Phone,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		r.logger.Error(logger.LogInput{
			Action: "PostgresUserRepository.Save",
			State:  logger.FAILED,
			Error: &logger.Error{
				Message: err.Error(),
			},
		})
		return errors.New("Internal server error")
	}

	return nil
}
