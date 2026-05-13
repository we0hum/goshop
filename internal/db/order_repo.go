package db

import (
	"context"
	"goshop/internal/models"

	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Create(ctx context.Context, o *models.Order) (models.Order, error) {
	var ord models.Order

	query := `
	INSERT INTO orders (product_id, quantity, total)
	VALUES ($1, $2, $3)
	RETURNING id, product_id, quantity, total;
	`

	err := r.db.GetContext(
		ctx,
		&ord,
		query,
		o.ProductID,
		o.Quantity,
		o.Total,
	)

	if err != nil {
		return models.Order{}, err
	}

	return ord, nil
}
