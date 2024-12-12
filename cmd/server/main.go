package main

import (
	"log"
	"pt-xyz-multifinance/configs"
	"pt-xyz-multifinance/internal/handler"
	"pt-xyz-multifinance/internal/repository"
	"pt-xyz-multifinance/internal/service"
	"pt-xyz-multifinance/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	// Inisialisasi konfigurasi
	config := configs.LoadConfig()

	// Koneksi database
	db := configs.InitDatabase(config)

	// Inisialisasi repository
	consumerRepo := repository.NewConsumerRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	limitRepo := repository.NewLimitRepository(db)

	// Inisialisasi service
	consumerService := service.NewConsumerService(consumerRepo)
	transactionService := service.NewTransactionService(transactionRepo, consumerRepo)
	limitService := service.NewLimitService(limitRepo)

	// Inisialisasi handler
	consumerHandler := handler.NewConsumerHandler(consumerService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	limitHandler := handler.NewLimitHandler(limitService)

	// Inisialisasi Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.RateLimiterMiddleware())
	e.Use(middleware.SecurityMiddleware())

	// Rute
	v1 := e.Group("/v1")
	{
		v1.POST("/consumers", consumerHandler.CreateConsumer)
		v1.GET("/consumers/:id", consumerHandler.GetConsumer)

		v1.POST("/transactions", transactionHandler.CreateTransaction)
		v1.GET("/transactions/:id", transactionHandler.GetTransaction)

		v1.GET("/limits/:consumerId", limitHandler.GetConsumerLimit)
		v1.POST("/limits/:consumerId", limitHandler.CreateConsumerLimits)
	}

	// Jalankan server
	log.Fatal(e.Start(":8080"))
}
