package services

import (
	"CRUD-Go-Hexa-MongoDB/internal/domain/product"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mock Repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) FindAll() ([]product.Product, error) {
	args := m.Called()
	return args.Get(0).([]product.Product), args.Error(1)
}

func (m *MockRepository) FindByID(id uuid.UUID) (product.Product, error) {
	args := m.Called(id)
	return args.Get(0).(product.Product), args.Error(1)
}

func (m *MockRepository) Create(product product.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockRepository) Update(product product.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestFindAll(t *testing.T) {
	// Setup
	mockRepo := new(MockRepository)
	productService := NewProductService(mockRepo)

	t.Run("returns all products", func(t *testing.T) {
		mockProducts := []product.Product{
			{
				ID:    uuid.New(),
				Name:  "Product 1",
				Stock: 10,
			},
			{
				ID:    uuid.New(),
				Name:  "Product 2",
				Stock: 20,
			},
		}

		mockRepo.On("FindAll").Return(mockProducts, nil)

		products, err := productService.FindAll()
		assert.NoError(t, err)
		assert.Equal(t, mockProducts, products)
		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil //reset expectations after each test
	})

	t.Run("returns empty list when no products found", func(t *testing.T) {
		mockRepo.On("FindAll").Return([]product.Product{}, nil)

		products, err := productService.FindAll()
		assert.NoError(t, err)
		assert.Empty(t, products)
		mockRepo.AssertExpectations(t)
	})
}

func TestFindByID(t *testing.T) {
	// Setup
	mockRepo := new(MockRepository)
	productService := NewProductService(mockRepo)

	t.Run("returns product by id", func(t *testing.T) {
		mockProduct := product.Product{
			ID:    uuid.New(),
			Name:  "Product 1",
			Stock: 10,
		}

		mockRepo.On("FindByID", mockProduct.ID).Return(mockProduct, nil)

		product, err := productService.FindByID(mockProduct.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockProduct, product)
		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil //reset expectations after each test
	})

	t.Run("returns error when product not found", func(t *testing.T) {
		nonExistentID := uuid.New()
		mockRepo.On("FindByID", nonExistentID).Return(product.Product{}, nil)

		foundProduct, err := productService.FindByID(nonExistentID)
		assert.Error(t, mongo.ErrNoDocuments, err)
		assert.Equal(t, product.Product{}, foundProduct)
		mockRepo.AssertExpectations(t)
	})
}
