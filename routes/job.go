package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omjikush09/kiranaClub/controllers"
)

func RegisterJobRoutes(api fiber.Router) {

	api.Post("/submit", controllers.CreateJob)
	api.Get("/status", controllers.GetJob)
}
