package services

import (
	"context"
	"fmt"
	"goshop/internal/cache"
	"goshop/internal/db"
	"goshop/internal/models"
)

type ProductService struct {
	repo  *db.ProductRepo
	cache *cache.ProductCache
}

func NewProductService(repo *db.ProductRepo, cache *cache.ProductCache) *ProductService {
	return &ProductService{repo: repo, cache: cache}
}

func (s *ProductService) GetProduct(ctx context.Context, id int) (*models.Product, error) {
	// 1️⃣ Проверяем Redis
	if p, _ := s.cache.Get(ctx, id); p != nil {
		fmt.Println("Из Redis")
		return p, nil
	}

	// 2️⃣ Если нет в кэше — берём из БД
	fmt.Println("Из PostgreSQL")
	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 3️⃣ Сохраняем в Redis
	_ = s.cache.Set(ctx, p)
	return p, nil
}
