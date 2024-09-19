//

package product

import "github.com/google/uuid"

type Repository interface {
	FindAll() ([]Product, error)
	FindByID(id uuid.UUID) (Product, error)
	Create(product Product) error
	Update(product Product) error
	Delete(id uuid.UUID) error
}
