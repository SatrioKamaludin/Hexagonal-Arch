package ports

import (
	product "CRUD-Go-Hexa-MongoDB/internal/domain/models"

	"github.com/google/uuid"
)

// Repository defines the interface for product operations
type Repository interface {
	FindAll() ([]product.Product, error) // Ensure the correct product type
	FindByID(id uuid.UUID) (product.Product, error)
	Create(product product.Product) error // Use product.Product here
	Update(product product.Product) error
	Delete(id uuid.UUID) error
}
