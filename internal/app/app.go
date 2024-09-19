package app

import (
	"CRUD-Go-Hexa-MongoDB/internal/controllers"
	mongoRepo "CRUD-Go-Hexa-MongoDB/internal/repository/mongo"
	"CRUD-Go-Hexa-MongoDB/internal/services"
	"CRUD-Go-Hexa-MongoDB/pkg/config"

	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Setup() *fiber.App {
	cfg := config.LoadConfig()

	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	client, err := mongoDriver.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(cfg.DBName)
	productRepo := mongoRepo.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	app := fiber.New()
	app.Get("/products", productController.FindAll)
	app.Get("/products/:id", productController.FindByID)
	app.Post("/products", productController.Create)
	app.Put("/products/:id", productController.Update)
	app.Delete("/products/:id", productController.Delete)

	return app
}
