// act as a router

package app

import (
	handlers "CRUD-Go-Hexa-MongoDB/internal/handlers"
	"database/sql"
	"fmt"

	// mongoRepo "CRUD-Go-Hexa-MongoDB/internal/repository/mongo"
	postgreSQLRepo "CRUD-Go-Hexa-MongoDB/internal/repository/postgresql"
	"CRUD-Go-Hexa-MongoDB/internal/services"
	"CRUD-Go-Hexa-MongoDB/pkg/config"

	// "context"
	"log"
	// "time"

	"github.com/gofiber/fiber/v2"
	// mongoDriver "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	_ "github.com/lib/pq"
)

func Setup() *fiber.App {
	cfg := config.LoadConfig()

	// clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	// client, err := mongoDriver.Connect(context.Background(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPass, cfg.PostgresDBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	productRepo := postgreSQLRepo.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := handlers.NewProductController(productService)

	app := fiber.New()
	app.Get("/products", productController.FindAll)
	app.Get("/products/:id", productController.FindByID)
	app.Post("/products", productController.Create)
	app.Put("/products/:id", productController.Update)
	app.Delete("/products/:id", productController.Delete)

	return app
}
