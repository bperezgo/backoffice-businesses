package handlers

import (
	"github.com/bperezgo/backoffice-businesses/config"
	"github.com/bperezgo/backoffice-businesses/repositories"
	"github.com/bperezgo/backoffice-businesses/services"
	"github.com/bperezgo/backoffice-businesses/shared/platform/handler"
	"github.com/bperezgo/backoffice-businesses/shared/platform/repositories/postgres"
)

func InitHandlers(c *config.Config) []handler.Handler {
	getHealth := NewGetHealthHandler()
	createOrders := NewCreateOrderHandler()

	postgresConnection := postgres.NewPostgresPoolConnection(postgres.Options{
		Host:         c.POSTGRES_HOST,
		Port:         c.POSTGRES_PORT,
		Username:     c.POSTGRES_USERNAME,
		Password:     c.POSTGRES_PASSWORD,
		Database:     c.POSTGRES_DATABASE,
		SSLMode:      c.POSTGRES_SSLMODE,
		MaxIdleConns: c.POSTGRES_MAX_IDLE_CONNS,
		MaxOpenConns: c.POSTGRES_MAX_OPEN_CONNS,
	})

	usersRepository := repositories.NewPostgresUserRepository(postgresConnection.DB)
	usersService := services.NewUserService(usersRepository)
	createUsers := NewCreateUserHandler(usersService)

	return []handler.Handler{
		getHealth,
		createOrders,
		createUsers,
	}
}
