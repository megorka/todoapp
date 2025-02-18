package service

import (
	"github.com/megorka/todoapp/user_service/models"
	"github.com/megorka/todoapp/user_service/repository"
)

type Service struct {
	repo *repository.Repository
}


func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(username, email, password string) error {
	return s.repo.CreateUser(username, email, password)
}

func (s *Service) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}

// func (s *Service) HandleUserCreation(ctx context.Context) {
// 	reader := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers: []string{"localhost:9092"},
// 		Topic: "user-creation",
// 		GroupID: "user-service-group",
// 	})

// 	defer reader.Close()

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Stopping Kafka consumer...")
// 			return
// 		default:
// 			msg, err := reader.ReadMessage(ctx)
// 			if err != nil {
// 				log.Printf("Error reading message: %v", err)
// 				continue
// 			}

// 			var event events.UserCreatedEvent

// 			err = json.Unmarshal(msg.Value, &event)
// 			if err != nil {
// 				log.Printf("Failed to unmarshal message: %v", err)
// 				continue
// 			}

// 			err = s.repo.CreateUser(event.Username, event.Email, event.Password)
// 			if err != nil {
// 				log.Printf("Failed to create user: %v", err)
// 				continue
// 			}
// 			log.Printf("User created: %s", event.Email)
// 		}
// 	}
// }