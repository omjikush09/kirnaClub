package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omjikush09/kiranaClub/database"
	"github.com/omjikush09/kiranaClub/models"
	"github.com/omjikush09/kiranaClub/utils"
)

func CreateStore(c *fiber.Ctx) error {
	var store models.Store

	if err := c.BodyParser(&store); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if validationErrors, isValid := utils.ValidateStruct(store); !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationErrors,
		})
	}

	if err := database.DB.Create(&store).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    store,
	})

}
