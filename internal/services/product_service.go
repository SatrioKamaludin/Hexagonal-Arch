// Contains core business logic and interacts with Repository interface
package services

import (
	"CRUD-Go-Hexa-MongoDB/internal/domain/product"

	"github.com/google/uuid"
)

type ProductService struct {
	productRepo product.Repository
}

func NewProductService(productRepo product.Repository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (s *ProductService) FindAll() ([]product.Product, error) {
	return s.productRepo.FindAll()
}

func (s *ProductService) FindByID(id uuid.UUID) (product.Product, error) {
	return s.productRepo.FindByID(id)
}

func (s *ProductService) Create(product product.Product) error {
	return s.productRepo.Create(product)
}

func (s *ProductService) Update(product product.Product) error {
	return s.productRepo.Update(product)
}

func (s *ProductService) Delete(id uuid.UUID) error {
	return s.productRepo.Delete(id)
}
