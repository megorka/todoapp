package config

import (
	"os"

	"github.com/megorka/todoapp/user_service/pkg/postgres"
	router "github.com/megorka/todoapp/user_service/transport/http"
)

type Config struct {
	DB postgres.DBConfig
	RouterConfig router.Config
}

func NewConfig() *Config {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "megorka"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "todoapp"
	}

	return &Config{
		DB: postgres.DBConfig{
			Host: dbHost,
			Port: dbPort,
			User: dbUser,
			Password: dbPassword,
			Database: dbName,
		},
		RouterConfig: router.Config{
			Host: "localhost",
			Port: "8181",
		},
	}
}