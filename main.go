package main

import (
	"HexAndClean/adapters"
	"HexAndClean/core"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	app := fiber.New()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	orderRepository := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepository)
	orderHandler := adapters.NewHttpOrderHandler(orderService)
	app.Post("/orders", orderHandler.CreateOrder)

	db.AutoMigrate(&core.Order{})

	app.Listen(":8000")
}
