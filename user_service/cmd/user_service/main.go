package main

import (
	"log"

	"github.com/megorka/todoapp/user_service/config"
	"github.com/megorka/todoapp/user_service/pkg/postgres"
	"github.com/megorka/todoapp/user_service/repository"
	"github.com/megorka/todoapp/user_service/service"
	router "github.com/megorka/todoapp/user_service/transport/http"
)

func main() {

	cfg := config.NewConfig()

	db, err := postgres.NewPostgres(cfg.DB)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	repository := repository.NewRepository(db)

	if err := repository.CreateUsers(); err != nil {
		log.Fatalf("CreateUsers: %v", err)
	}

	service := service.NewService(repository)

	// go func(){
	// 	log.Println("Starting Kafka consumer...")
	// 	ctx := context.Background()
	// 	service.HandleUserCreation(ctx)
	// }()

	handler := router.NewHandler(service)

	r := router.NewRouter(cfg.RouterConfig, handler)
	
	r.Run()
}