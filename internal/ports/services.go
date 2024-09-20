package ports

import (
	"CRUD-Go-Hexa-MongoDB/internal/domain/models"
	"CRUD-Go-Hexa-MongoDB/internal/utils"
)

type IProductService interface {
	FindAll() utils.ServiceResponse
	FindByID(idStr string) utils.ServiceResponse
	Create(productData map[string]string) utils.ServiceResponse
	Update(idStr string, productData map[string]string) utils.ServiceResponse
	Delete(idStr string) utils.ServiceResponse
}

type IProfilingService interface {
	Log(profiling models.Profiling) error
}
