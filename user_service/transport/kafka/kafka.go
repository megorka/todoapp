package kafka

import (
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
