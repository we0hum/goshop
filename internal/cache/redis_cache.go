package cache

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"goshop/internal/models"

	"github.com/redis/go-redis/v9"
)

type ProductCache struct {
	rdb *redis.Client
}

func NewProductCache(addr string) *ProductCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &ProductCache{rdb: rdb}
}

func (c *ProductCache) Get(ctx context.Context, id int) (*models.Product, error) {
	val, err := c.rdb.Get(ctx, key(id)).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var p models.Product
	if err := json.Unmarshal([]byte(val), &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (c *ProductCache) Set(ctx context.Context, p *models.Product) error {
	data, _ := json.Marshal(p)
	return c.rdb.Set(ctx, key(p.ID), data, 5*time.Minute).Err()
}

func key(id int) string {
	return "product:" + strconv.Itoa(id)
}
