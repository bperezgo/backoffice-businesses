package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Options struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SSLMode  string

	MaxIdleConns int
	MaxOpenConns int
}

type PostgresPoolConnection struct {
	DB *sqlx.DB
}

func NewPostgresPoolConnection(opts Options) *PostgresPoolConnection {
	driverName := "postgres"
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		opts.Host,
		opts.Port,
		opts.Username,
		opts.Password,
		opts.Database,
		opts.SSLMode,
	)

	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	return &PostgresPoolConnection{
		DB: db,
	}
}
