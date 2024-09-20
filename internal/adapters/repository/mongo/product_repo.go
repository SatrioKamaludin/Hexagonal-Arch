package mongo

import (
	product "CRUD-Go-Hexa-MongoDB/internal/domain/models" // Import only product entity
	"CRUD-Go-Hexa-MongoDB/internal/ports"                 // Import ports for Repository interface
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

// NewProductRepository now returns ports.Repository instead of product.Repository
func NewProductRepository(db *mongo.Database) ports.IProductRepository {
	return &ProductRepository{
		collection: db.Collection("products"),
	}
}

func (r *ProductRepository) FindAll() ([]product.Product, error) {
	var products []product.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product product.Product
		cursor.Decode(&product)
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) FindByID(id uuid.UUID) (product.Product, error) {
	var product product.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&product)
	return product, err
}

func (r *ProductRepository) Create(product product.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, product)
	return err
}

func (r *ProductRepository) Update(product product.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"id": product.ID}, bson.M{"$set": product})
	return err
}

func (r *ProductRepository) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"id": id})
	return err
}
