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

func (r *OrderRepo) Create(ctx context.Context, o *models.Order) (models.Order, error) {
	var ord models.Order

	query := `
 INSERT INTO orders (product_id, quantity, total)
 VALUES ($1, $2, $3)
 RETURNING id, product_id, quantity, total;
 `

	if err := r.db.GetContext(ctx, &ord, query,
		o.ProductID, o.Quantity, o.Total); err != nil {
		return models.Order{}, err
	}
	return ord, nil
}
