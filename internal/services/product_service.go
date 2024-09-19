// Contains core business logic and interacts with Repository interface
package services

import (
	"CRUD-Go-Hexa-MongoDB/internal/domain/product"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
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
	if product.Name == "" {
		return errors.New("product name cannot be empty") // for service-side unit test Create
	}
	if product.Stock < 0 {
		return errors.New("product stock cannot be negative") // for service-side unit test Create
	}
	return s.productRepo.Create(product)
}

func (s *ProductService) Update(product product.Product) error {
	if product.Stock < 0 {
		return errors.New("product stock cannot be negative") // for service-side unit test Update
	}

	err := s.productRepo.Update(product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("product not found")
		}
		return err
	}

	return s.productRepo.Update(product)
}

func (s *ProductService) Delete(id uuid.UUID) error {
	return s.productRepo.Delete(id)
}
