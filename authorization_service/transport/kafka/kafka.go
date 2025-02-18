package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/megorka/todoapp/authorization/events"
	"github.com/megorka/todoapp/authorization/pkg/jwt"
	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	return &KafkaProducer{writer: &kafka.Writer{
		Addr: kafka.TCP(brokers...),
		Topic: topic,
		Balancer: &kafka.LeastBytes{},
	}}
}

func (p *KafkaProducer) SendUserCreatedEvent(username, email, password string) error {

	hashedPassword, err := jwt.HashPassword(password)
	if err != nil {
		return fmt.Errorf("HashPassword: %w", err)
	}

	event := events.UserCreatedEvent{
		Username: username,
		Email: email,
		Password: hashedPassword,
	}

	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("JsonMarshal: %w", err)
	}
	
	err = p.writer.WriteMessages(context.Background(), kafka.Message{Value: payload})

	if err != nil {
		return fmt.Errorf("WriteMessages: %w", err)
	}

	log.Println("User creation event sent to kafka")
	return nil
}