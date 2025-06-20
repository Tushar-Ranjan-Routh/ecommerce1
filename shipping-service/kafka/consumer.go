package kafka

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

func StartConsumer() {
	log.Println("🚀 Starting shipping service Kafka consumer...")

	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "kafka:9092"
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    "order-events",
		GroupID:  "shipping-service",
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	defer r.Close()

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		m, err := r.FetchMessage(ctx)
		cancel()

		if err == nil {
			log.Printf("✅ Consumed message: %s\n", string(m.Value))
			if err := r.CommitMessages(context.Background(), m); err != nil {
				log.Printf("⚠️ Failed to commit message: %v", err)
			}
			continue
		}

		log.Printf("❌ Kafka not ready or error fetching message: %v\n", err)
		time.Sleep(5 * time.Second)
	}
}
