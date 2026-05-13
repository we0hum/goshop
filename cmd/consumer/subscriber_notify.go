package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "orders",
		GroupID: "notifications-group",
	})

	fmt.Println("👂 NotificationService слушает Kafka topic 'orders'...")

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			continue
		}
		fmt.Println("📦 Новое событие:", string(msg.Value))
	}
}
