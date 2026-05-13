package services

import (
	"context"
	"goshop/internal/db"
	"goshop/internal/models"
)

type OrderService struct {
	repo *db.OrderRepo
}

func NewOrderService(repo *db.OrderRepo) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, o *models.Order) error {
	return s.repo.Create(ctx, o)
}
