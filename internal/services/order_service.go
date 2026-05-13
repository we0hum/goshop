package services

import (
	"context"
	"goshop/internal/db"
	"goshop/internal/events"
	"goshop/internal/models"
)

type OrderService struct {
	repo      *db.OrderRepo
	publisher *events.Publisher
}

func NewOrderService(repo *db.OrderRepo, publisher *events.Publisher) *OrderService {
	return &OrderService{repo: repo, publisher: publisher}
}

func (s *OrderService) CreateOrder(ctx context.Context, o *models.Order) (models.Order, error) {
	order, err := s.repo.Create(ctx, o)
	if err != nil {
		return models.Order{}, err
	}

	_ = s.publisher.PublishOrderCreated(ctx, order.ID)

	return order, nil
}
