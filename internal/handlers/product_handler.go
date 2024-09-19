// Handles HTTP requests and responses. Acts as an entry point to the app, calling service layer to do business logic

package controllers

import (
	"CRUD-Go-Hexa-MongoDB/internal/ports"

	"github.com/gofiber/fiber/v2"
)

type ProdctHandler struct {
	productService ports.IProductService
}

func NewProductController(productService ports.IProductService) *ProdctHandler {
	return &ProdctHandler{
		productService: productService,
	}
}

func (c *ProdctHandler) FindAll(ctx *fiber.Ctx) error {
	response := c.productService.FindAll()
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProdctHandler) FindByID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	response := c.productService.FindByID(idStr)
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProdctHandler) Create(ctx *fiber.Ctx) error {
	productData := map[string]string{
		"name":  ctx.FormValue("name"),
		"stock": ctx.FormValue("stock"),
	}

	response := c.productService.Create(productData)
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProdctHandler) Update(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	productData := map[string]string{
		"name":  ctx.FormValue("name"),
		"stock": ctx.FormValue("stock"),
	}

	response := c.productService.Update(idStr, productData)
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProdctHandler) Delete(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	response := c.productService.Delete(idStr)
	return ctx.Status(response.Code).JSON(response)
}
