// act as a router

package app

import (
	handlers "CRUD-Go-Hexa-MongoDB/internal/adapters/handlers"
	"context"
	"database/sql"
	"fmt"

	mongoRepo "CRUD-Go-Hexa-MongoDB/internal/adapters/repository/mongo"
	postgreSQLRepo "CRUD-Go-Hexa-MongoDB/internal/adapters/repository/postgresql"
	"CRUD-Go-Hexa-MongoDB/internal/domain/services"
	"CRUD-Go-Hexa-MongoDB/pkg/config"

	// "context"
	"log"
	// "time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Setup() *fiber.App {
	cfg := config.LoadConfig()

	//Connect to MongoDB
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongoDriver.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//Ping to MongoDB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	//Obtain reference to MongoDB database
	mongoDB := client.Database(cfg.DBName)

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//Connect to PostgreSQL
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

	profilingRepo := mongoRepo.NewProfilingRepository(mongoDB)
	profilingService := services.NewProfilingService(profilingRepo)

	productRepo := postgreSQLRepo.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := handlers.NewProductController(productService, profilingService)

	app := fiber.New()
	app.Get("/products", productController.FindAll)
	app.Get("/products/:id", productController.FindByID)
	app.Post("/products", productController.Create)
	app.Put("/products/:id", productController.Update)
	app.Delete("/products/:id", productController.Delete)

	return app
}
