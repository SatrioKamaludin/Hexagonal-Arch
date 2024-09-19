// Contains core business logic and interacts with Repository interface
package services

import (
	product "CRUD-Go-Hexa-MongoDB/internal/domain/models"
	"CRUD-Go-Hexa-MongoDB/internal/ports"
	"CRUD-Go-Hexa-MongoDB/internal/utils"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	productRepo ports.Repository
}

func NewProductService(productRepo ports.Repository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (s *ProductService) FindAll() utils.ServiceResponse {
	products, err := s.productRepo.FindAll()
	if err != nil {
		return utils.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch products",
			Data:    err.Error(),
		}
	}
	return utils.ServiceResponse{
		Code:    http.StatusOK,
		Message: "Products fetched successfully",
		Data:    products,
	}
}

func (s *ProductService) FindByID(idStr string) utils.ServiceResponse {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return utils.ServiceResponse{
			Code:    http.StatusNotFound,
			Message: "Product with ID " + idStr + " not found",
			Data:    nil,
		}
	}

	product, err := s.productRepo.FindByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return utils.ServiceResponse{
				Code:    http.StatusNotFound,
				Message: "Product with ID " + idStr + " not found",
				Data:    nil,
			}
		}
		return utils.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch product",
			Data:    err.Error(),
		}
	}
	return utils.ServiceResponse{
		Code:    http.StatusOK,
		Message: "Product fetched successfully",
		Data:    product,
	}
}

func (s *ProductService) Create(productData map[string]string) utils.ServiceResponse {
	name := productData["name"]
	stockStr := productData["stock"]

	// Validate inputs
	var arrErrors []string

	if name == "" {
		arrErrors = append(arrErrors, "Name cannot be empty")
	}

	stock, err := strconv.Atoi(stockStr)
	if err != nil || stock < 0 {
		arrErrors = append(arrErrors, "Invalid Stock Value, must be a number and greater than 0")
	}

	if len(arrErrors) > 0 {
		return utils.ServiceResponse{
			Code:    http.StatusBadRequest,
			Message: "Validation error",
			Data:    arrErrors,
		}
	}

	// Create the product
	product := product.Product{
		ID:    uuid.New(),
		Name:  name,
		Stock: stock,
	}

	err = s.productRepo.Create(product)
	if err != nil {
		return utils.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error creating product",
			Data:    err.Error(),
		}
	}

	return utils.ServiceResponse{
		Code:    http.StatusCreated,
		Message: "Product created successfully",
		Data:    product,
	}
}

func (s *ProductService) Update(idStr string, productData map[string]string) utils.ServiceResponse {
	// Parse the product ID. If invalid, treat it as "not found"
	id, err := uuid.Parse(idStr)
	if err != nil {
		return utils.ServiceResponse{
			Code:    http.StatusNotFound,
			Message: "Product with ID " + idStr + " not found",
			Data:    nil,
		}
	}

	existingProduct, err := s.productRepo.FindByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return utils.ServiceResponse{
				Code:    http.StatusNotFound,
				Message: "Product with ID " + id.String() + " not found",
				Data:    nil,
			}
		}
		return utils.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch product",
			Data:    err.Error(),
		}
	}

	name := productData["name"]
	stockStr := productData["stock"]

	if name != "" {
		existingProduct.Name = name
	}

	if stockStr != "" {
		stock, err := strconv.Atoi(stockStr)
		if err != nil || stock < 0 {
			return utils.ServiceResponse{
				Code:    http.StatusBadRequest,
				Message: "Invalid Stock Value, must be a number and greater than 0",
				Data:    nil,
			}
		}
		existingProduct.Stock = stock
	}

	err = s.productRepo.Update(existingProduct)
	if err != nil {
		return utils.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error updating product",
			Data:    err.Error(),
		}
	}

	return utils.ServiceResponse{
		Code:    http.StatusOK,
		Message: "Product updated successfully",
		Data:    existingProduct,
	}
}

func (s *ProductService) Delete(idStr string) utils.ServiceResponse {
	// Parse the product ID. If invalid, treat it as "not found"
	id, err := uuid.Parse(idStr)
	if err != nil {
		return utils.ServiceResponse{
			Code:    http.StatusNotFound,
			Message: "Product with ID " + idStr + " not found",
			Data:    nil,
		}
	}

	// Find and delete the product by ID
	err = s.productRepo.Delete(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return utils.ServiceResponse{
				Code:    http.StatusNotFound,
				Message: "Product with ID " + idStr + " not found",
				Data:    nil,
			}
		}
		return utils.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error deleting product",
			Data:    err.Error(),
		}
	}

	// Return success response
	return utils.ServiceResponse{
		Code:    http.StatusOK,
		Message: "Product deleted successfully",
		Data:    nil,
	}
}
