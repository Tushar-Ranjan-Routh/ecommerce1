package kafka

import (
	"context"
	"encoding/json"
	"log"
	"order-service/models"
	"os"

	"github.com/segmentio/kafka-go"
)

var Writer *kafka.Writer

func InitKafka() {
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "kafka:9092"
	}
	Writer = &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    "order-events",
		Balancer: &kafka.LeastBytes{},
	}
	log.Println("✅ Kafka producer initialized at", broker)
}

func PublishOrder(order models.Order) {
	data, _ := json.Marshal(order)
	msg := kafka.Message{
		Key:   []byte(order.ID),
		Value: data,
	}

	if err := Writer.WriteMessages(context.Background(), msg); err != nil {
		log.Println("❌ Failed to publish to Kafka:", err)
	}
}
