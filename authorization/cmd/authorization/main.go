package authorization

import (
	"os"

	"github.com/megorka/todoapp/authorization/config"
	"github.com/megorka/todoapp/authorization/service"
	router "github.com/megorka/todoapp/authorization/transport/http"
	"github.com/megorka/todoapp/authorization/transport/kafka"
)

func main() {

	cfg := config.NewConfig()

	kafkaProducer := kafka.NewKafkaProducer([]string{os.Getenv("Kafkahost")}, "user-creation")

	service := service.NewService(kafkaProducer)

	handler := router.NewHandler(service)

	r := router.NewRouter(cfg.RouterConfig, handler)
	
	r.Run()
}