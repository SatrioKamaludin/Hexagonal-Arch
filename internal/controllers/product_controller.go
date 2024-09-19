// Handles HTTP requests and responses. Acts as an entry point to the app, calling service layer to do business logic

package controllers

import (
	"CRUD-Go-Hexa-MongoDB/internal/domain/product"
	"CRUD-Go-Hexa-MongoDB/internal/services"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (c *ProductController) FindAll(ctx *fiber.Ctx) error {
	products, err := c.productService.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(products)
}

func (c *ProductController) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	product, err := c.productService.FindByID(uuid.MustParse(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(product)
}

func (c *ProductController) Create(ctx *fiber.Ctx) error {
	name := ctx.FormValue("name")
	stockStr := ctx.FormValue("stock")

	if name == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name cannot be empty"})
	}

	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Invalid Stock Value": err.Error(),
		})
	}

	product := product.Product{
		ID:    uuid.New(),
		Name:  name,
		Stock: stock,
	}

	if err := c.productService.Create(product); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(product)
}

func (c *ProductController) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Invalid ID": err.Error(),
		})
	}

	name := ctx.FormValue("name")
	stockStr := ctx.FormValue("stock")

	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Invalid Stock Value": err.Error(),
		})
	}

	product := product.Product{
		ID:    id,
		Name:  name,
		Stock: stock,
	}
	return ctx.Status(fiber.StatusOK).JSON(product)
}

func (c *ProductController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.productService.Delete(uuid.MustParse(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.SendStatus(fiber.StatusOK)
}
