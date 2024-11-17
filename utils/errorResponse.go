package utils

import "github.com/gofiber/fiber/v2"

func InternalServerError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Something went wrong",
	})
}
