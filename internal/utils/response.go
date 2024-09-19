package utils

import "github.com/gofiber/fiber/v2"

func NotFoundResponse(c *fiber.Ctx, idStr string) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Product with ID " + idStr + " not found",
	})
}

func InternalServerErrorResponse(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": err.Error(),
	})
}
