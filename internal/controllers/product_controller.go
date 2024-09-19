// Handles HTTP requests and responses. Acts as an entry point to the app, calling service layer to do business logic

package controllers

import (
	"CRUD-Go-Hexa-MongoDB/internal/domain/product"
	"CRUD-Go-Hexa-MongoDB/internal/services"
	"CRUD-Go-Hexa-MongoDB/internal/utils"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
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
		return utils.InternalServerErrorResponse(ctx, err)
	}
	return ctx.Status(fiber.StatusOK).JSON(products)
}

func (c *ProductController) FindByID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := uuid.Parse(idStr)

	product, err := c.productService.FindByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return utils.NotFoundResponse(ctx, idStr)
		}
		return utils.InternalServerErrorResponse(ctx, err)
	}
	return ctx.Status(fiber.StatusOK).JSON(product)
}

func (c *ProductController) Create(ctx *fiber.Ctx) error {
	name := ctx.FormValue("name")
	stockStr := ctx.FormValue("stock")

	var errCount = 0
	var arrErrors = make([]string, 0)

	if name == "" {
		errCount++
		arrErrors = append(arrErrors, "Name cannot be empty")
	}

	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		errCount++
		arrErrors = append(arrErrors, "Invalid Stock Value, must be a number and not empty")
	}

	if errCount > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": arrErrors,
		})
	}

	product := product.Product{
		ID:    uuid.New(),
		Name:  name,
		Stock: stock,
	}

	if err := c.productService.Create(product); err != nil {
		return utils.InternalServerErrorResponse(ctx, err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(product)
}

func (c *ProductController) Update(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := uuid.Parse(idStr)

	existingProduct, err := c.productService.FindByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return utils.NotFoundResponse(ctx, idStr)
		}
		return utils.InternalServerErrorResponse(ctx, err)
	}

	name := ctx.FormValue("name")
	stockStr := ctx.FormValue("stock")

	if name != "" {
		existingProduct.Name = name
	}

	if stockStr != "" {
		stock, err := strconv.Atoi(stockStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"Invalid Stock Value": err.Error(),
			})
		}
		existingProduct.Stock = stock
	}

	if err := c.productService.Update(existingProduct); err != nil {
		return utils.InternalServerErrorResponse(ctx, err)
	}
	return ctx.Status(fiber.StatusOK).JSON(existingProduct)
}

func (c *ProductController) Delete(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := uuid.Parse(idStr)

	// Check if the product exists
	_, err = c.productService.FindByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return utils.NotFoundResponse(ctx, idStr)
		}
		return utils.InternalServerErrorResponse(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
