package services

import (
	product "CRUD-Go-Hexa-MongoDB/internal/domain/models"
	"net/http"
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

		response := productService.FindAll()
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Products fetched successfully", response.Message)
		assert.Equal(t, mockProducts, response.Data)
		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil //reset expectations after each test
	})

	t.Run("returns empty list when no products found", func(t *testing.T) {
		mockRepo.On("FindAll").Return([]product.Product{}, nil)

		response := productService.FindAll()
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Products fetched successfully", response.Message)
		assert.Empty(t, response.Data)
		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil //reset expectations after each test
	})
}

func TestFindByID(t *testing.T) {
	// Setup
	mockRepo := new(MockRepository)
	productService := NewProductService(mockRepo)

	id := uuid.New()

	t.Run("returns product by id", func(t *testing.T) {
		mockProduct := product.Product{
			ID:    id,
			Name:  "Product 1",
			Stock: 10,
		}

		mockRepo.On("FindByID", id).Return(mockProduct, nil)

		response := productService.FindByID(id.String())
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Product fetched successfully", response.Message)
		assert.Equal(t, mockProduct, response.Data)
		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil //reset expectations after each test
	})

	t.Run("returns not found error when product not found", func(t *testing.T) {
		mockRepo.On("FindByID", id).Return(product.Product{}, mongo.ErrNoDocuments)

		response := productService.FindByID(id.String())
		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "Product with ID "+id.String()+" not found", response.Message)
		assert.Nil(t, response.Data)
		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil //reset expectations after each test
	})
}

func TestCreate(t *testing.T) {
	// Setup
	mockRepo := new(MockRepository)
	productService := NewProductService(mockRepo)

	t.Run("creates product successfully", func(t *testing.T) {
		mockProduct := product.Product{
			Name:  "Product 1",
			Stock: 10,
		}

		productData := map[string]string{
			"name":  "Product 1",
			"stock": "10",
		}

		// Simulate repository behavior with the fixed UUID
		mockRepo.On("Create", mock.MatchedBy(func(p product.Product) bool {
			return p.Name == mockProduct.Name && p.Stock == mockProduct.Stock
		})).Return(nil)

		response := productService.Create(productData)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Equal(t, "Product created successfully", response.Message)

		// Assert that the response.Data matches the expected product but ignore the ID
		actualProduct, ok := response.Data.(product.Product)
		assert.True(t, ok)
		assert.Equal(t, mockProduct.Name, actualProduct.Name)
		assert.Equal(t, mockProduct.Stock, actualProduct.Stock)

		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil // reset expectations after each test
	})

	t.Run("returns validation error when name is empty", func(t *testing.T) {
		productData := map[string]string{
			"name":  "",
			"stock": "10",
		}

		response := productService.Create(productData)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "Validation error", response.Message)
		assert.Contains(t, response.Data, "Name cannot be empty")
		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil // reset expectations after each test
	})

	t.Run("returns validation error when stock is invalid", func(t *testing.T) {
		productData := map[string]string{
			"name":  "Product 1",
			"stock": "-1",
		}

		response := productService.Create(productData)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "Validation error", response.Message)
		assert.Contains(t, response.Data, "Invalid Stock Value, must be a number and greater than 0")
		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil // reset expectations after each test
	})
}

func TestUpdate(t *testing.T) {
	// Setup
	mockRepo := new(MockRepository)
	productService := NewProductService(mockRepo)

	t.Run("updates product successfully", func(t *testing.T) {
		id := uuid.New()
		mockProduct := product.Product{
			ID:    id,
			Name:  "Product 1",
			Stock: 10,
		}

		// Mock FindByID to return product that needs updating
		mockRepo.On("FindByID", id).Return(mockProduct, nil)

		// Mock Update to confirm that it's called using correct params
		mockRepo.On("Update", mockProduct).Return(nil)

		productData := map[string]string{
			"name":  "Product 1",
			"stock": "10",
		}

		response := productService.Update(id.String(), productData)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Product updated successfully", response.Message)
		assert.Equal(t, mockProduct, response.Data)
		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil //reset expectations after each test
	})

	t.Run("returns error when product not found", func(t *testing.T) {
		nonExistentID := uuid.New()

		productData := map[string]string{
			"name":  "Product 1",
			"stock": "10",
		}

		mockRepo.On("FindByID", nonExistentID).Return(product.Product{}, mongo.ErrNoDocuments)

		response := productService.Update(nonExistentID.String(), productData)
		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "Product with ID "+nonExistentID.String()+" not found", response.Message)
		assert.Nil(t, response.Data)
		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil //reset expectations after each test
	})

	t.Run("returns error when stock is negative", func(t *testing.T) {
		id := uuid.New()
		productData := map[string]string{
			"name":  "Product 1",
			"stock": "-1",
		}

		mockRepo.On("FindByID", id).Return(product.Product{}, nil)

		response := productService.Update(id.String(), productData)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "Invalid Stock Value, must be a number and greater than 0", response.Message)
		assert.Nil(t, response.Data)
		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil //reset expectations after each test
	})
}

func TestDelete(t *testing.T) {
	// Setup
	mockRepo := new(MockRepository)
	productService := NewProductService(mockRepo)

	t.Run("deletes product successfully", func(t *testing.T) {
		id := uuid.New()

		mockRepo.On("Delete", id).Return(nil)

		response := productService.Delete(id.String())
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Product deleted successfully", response.Message)
		assert.Nil(t, response.Data)
		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil //reset expectations after each test
	})

	t.Run("returns error when product not found", func(t *testing.T) {
		nonExistentID := uuid.New()
		mockRepo.On("Delete", nonExistentID).Return(mongo.ErrNoDocuments)

		response := productService.Delete(nonExistentID.String())
		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "Product with ID "+nonExistentID.String()+" not found", response.Message)
		assert.Nil(t, response.Data)
		mockRepo.AssertExpectations(t)
		mockRepo.ExpectedCalls = nil //reset expectations after each test
	})
}
