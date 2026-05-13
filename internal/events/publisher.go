package events

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Publisher struct {
	writer *kafka.Writer
}

func NewPublisher(broker, topic string) *Publisher {
	return &Publisher{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(broker),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *Publisher) PublishOrderCreated(ctx context.Context, orderID int) error {
	msg := fmt.Sprintf("Создан новый заказ #%d", orderID)
	return p.writer.WriteMessages(ctx, kafka.Message{
		Value: []byte(msg),
	})
}
