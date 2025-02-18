package authorization

import (
	"log"
	"os"

	"github.com/megorka/todoapp/authorization/config"
	"github.com/megorka/todoapp/authorization/pkg/postgres"
	"github.com/megorka/todoapp/authorization/repository"
	"github.com/megorka/todoapp/authorization/service"
	router "github.com/megorka/todoapp/authorization/transport/http"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Ошибка открытия файла лога:", err)
	}
	logger := log.New(file, "BOOKSERVICE: ", log.Ldate|log.Ltime|log.Lshortfile)

	cfg := config.NewConfig()

	db, err := postgres.NewPostgres(cfg.DB)
	if err != nil {
		logger.Fatal("Ошибка подключения к базе данных:", err)
	}

	repo := repository.NewRepository(db)


	service := service.NewService(repo)

	handler := router.NewHandler(service)

	r := router.NewRouter(cfg.RouterConfig, handler)
	
	r.Run()
}