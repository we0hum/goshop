package main

import (
	"goshop/internal/cache"
	"goshop/internal/db"
	"goshop/internal/events"
	"goshop/internal/handlers"
	"goshop/internal/services"
	"log"
	"net/http"
)

func main() {
	dsn := "postgres://postgres:password@localhost:5434/goshop?sslmode=disable"

	dbConn, err := db.NewPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}

	redisCache := cache.NewProductCache("localhost:6379")
	publisher := events.NewPublisher("localhost:9092", "orders")

	productRepo := db.NewProductRepo(dbConn)
	orderRepo := db.NewOrderRepo(dbConn)

	productService := services.NewProductService(productRepo, redisCache)
	orderService := services.NewOrderService(orderRepo, publisher)

	productHandler := handlers.NewProductHandler(productService)
	orderHandler := handlers.NewOrderHandler(orderService)

	http.HandleFunc("/products/", productHandler.GetProduct)
	http.HandleFunc("/orders", orderHandler.CreateOrder)

	log.Println("🚀 Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
