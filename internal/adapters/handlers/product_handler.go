// Handles HTTP requests and responses. Acts as an entry point to the app, calling service layer to do business logic

package handlers

import (
	"CRUD-Go-Hexa-MongoDB/internal/domain/models"
	"CRUD-Go-Hexa-MongoDB/internal/ports"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProdctHandler struct {
	productService   ports.IProductService
	profilingService ports.IProfilingService
}

func NewProductController(productService ports.IProductService, profilingService ports.IProfilingService) *ProdctHandler {
	return &ProdctHandler{
		productService:   productService,
		profilingService: profilingService,
	}
}
func (c *ProdctHandler) logProfiling(apiCall string, startTime time.Time) error {
	duration := time.Since(startTime).Milliseconds()

	profiling := models.Profiling{
		ID:        uuid.New(),
		APICall:   apiCall,
		Duration:  duration,
		Timestamp: time.Now(),
	}

	c.profilingService.Log(profiling)

	return nil
}

func (c *ProdctHandler) FindAll(ctx *fiber.Ctx) error {
	startTime := time.Now()
	response := c.productService.FindAll()
	c.logProfiling("FindAll", startTime)
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProdctHandler) FindByID(ctx *fiber.Ctx) error {
	startTime := time.Now()
	idStr := ctx.Params("id")
	response := c.productService.FindByID(idStr)
	c.logProfiling("FindByID: "+idStr, startTime)
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProdctHandler) Create(ctx *fiber.Ctx) error {
	startTime := time.Now()
	productData := map[string]string{
		"name":  ctx.FormValue("name"),
		"stock": ctx.FormValue("stock"),
	}

	response := c.productService.Create(productData)
	c.logProfiling("Create", startTime)
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProdctHandler) Update(ctx *fiber.Ctx) error {
	startTime := time.Now()
	idStr := ctx.Params("id")

	productData := map[string]string{
		"name":  ctx.FormValue("name"),
		"stock": ctx.FormValue("stock"),
	}

	response := c.productService.Update(idStr, productData)
	c.logProfiling("Update :"+idStr, startTime)
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProdctHandler) Delete(ctx *fiber.Ctx) error {
	startTime := time.Now()
	idStr := ctx.Params("id")

	response := c.productService.Delete(idStr)
	c.logProfiling("Delete: "+idStr, startTime)
	return ctx.Status(response.Code).JSON(response)
}
