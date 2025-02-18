package main

import (
	"github.com/megorka/todoapp/authorization/client"
	"github.com/megorka/todoapp/authorization/config"
	"github.com/megorka/todoapp/authorization/service"
	router "github.com/megorka/todoapp/authorization/transport/http"
)

func main() {

	cfg := config.NewConfig()

	userClient := client.NewClient("http://localhost:8181")

	service := service.NewService(userClient)

	handler := router.NewHandler(service)

	r := router.NewRouter(cfg.RouterConfig, *handler)
	
	r.Run()
}