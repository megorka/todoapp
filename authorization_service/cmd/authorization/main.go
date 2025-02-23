package main

import (
	"github.com/megorka/todoapp/authorization/adapters"
	"github.com/megorka/todoapp/authorization/config"
	"github.com/megorka/todoapp/authorization/service"
	router "github.com/megorka/todoapp/authorization/transport/http"
)

func main() {

	cfg := config.NewConfig()

	userClient := adapters.NewAdapter("http://localhost:8181")

	newService := service.NewService(userClient)

	handler := router.NewHandler(newService)

	r := router.NewRouter(cfg.RouterConfig, *handler)

	r.Run()
}
