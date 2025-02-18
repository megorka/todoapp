package config

import (
	router "github.com/megorka/todoapp/authorization/transport/http"
)

type Config struct {
	RouterConfig router.Config
}


func NewConfig() *Config {
	return &Config{
		RouterConfig: router.Config{
			Host: "localhost",
			Port: "8080",
		},
	}
}