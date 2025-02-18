package service

import (
	"context"
	"fmt"
	"log"

	"github.com/megorka/todoapp/authorization/transport/kafka"
)

type Service struct {
	kafka *kafka.KafkaProducer
}

func NewService(kafka *kafka.KafkaProducer) *Service {
	return &Service{kafka: kafka}
}

func (s *Service) RegisterUser(ctx context.Context, username, email, password string ) error {
	if err := s.kafka.SendUserCreatedEvent(username, email, password); err != nil {
		return fmt.Errorf("RegisterUser: %w", err)
	}
	log.Println("User registered: %s", email)
	return nil
}
