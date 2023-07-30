package config

import (
	"log"
	"os"
	"path"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort              int
	Environment             string
	POSTGRES_HOST           string
	POSTGRES_PORT           int
	POSTGRES_USERNAME       string
	POSTGRES_PASSWORD       string
	POSTGRES_DATABASE       string
	POSTGRES_SSLMODE        string
	POSTGRES_MAX_IDLE_CONNS int
	POSTGRES_MAX_OPEN_CONNS int
}

var config *Config

func InitConfig() error {
	environment, ok := os.LookupEnv("ENVIRONMENT")

	if !ok {
		return godotenv.Load()
	}

	if environment == "local" {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		p := path.Join(cwd, "../.env")
		return godotenv.Load(os.ExpandEnv(p))
	}

	return nil

}

func GetConfig() *Config {
	if config != nil {
		return config
	}

	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Panic("serverPort was not be processed")
	}

	environment := os.Getenv("ENVIRONMENT")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Panic("postgresPort was not be processed")
	}

	postgresUsername := os.Getenv("POSTGRES_USERNAME")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDatabase := os.Getenv("POSTGRES_DATABASE")
	postgresSSLMode := os.Getenv("POSTGRES_SSLMODE")
	postgresMaxIdleConns, err := strconv.Atoi(os.Getenv("POSTGRES_MAX_IDLE_CONNS"))

	if err != nil {
		log.Panic("postgresMaxIdleConns was not be processed")
	}
	postgresMaxOpenConns, err := strconv.Atoi(os.Getenv("POSTGRES_MAX_OPEN_CONNS"))

	if err != nil {
		log.Panic("postgresMaxOpenConns was not be processed")
	}

	config = &Config{
		ServerPort:              serverPort,
		Environment:             environment,
		POSTGRES_HOST:           postgresHost,
		POSTGRES_PORT:           postgresPort,
		POSTGRES_USERNAME:       postgresUsername,
		POSTGRES_PASSWORD:       postgresPassword,
		POSTGRES_DATABASE:       postgresDatabase,
		POSTGRES_SSLMODE:        postgresSSLMode,
		POSTGRES_MAX_IDLE_CONNS: postgresMaxIdleConns,
		POSTGRES_MAX_OPEN_CONNS: postgresMaxOpenConns,
	}
	return config
}
