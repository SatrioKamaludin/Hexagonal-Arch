package mongo

import (
	"CRUD-Go-Hexa-MongoDB/internal/domain/models"
	"CRUD-Go-Hexa-MongoDB/internal/ports"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProfilingRepository struct {
	collection *mongo.Collection
}

func NewProfilingRepository(db *mongo.Database) ports.IProfilingRepository {
	return &ProfilingRepository{
		collection: db.Collection("profiling"),
	}
}

func (r *ProfilingRepository) Create(profiling models.Profiling) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, profiling)
	return err
}
