package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/omjikush09/kiranaClub/config"
	"github.com/omjikush09/kiranaClub/cronJob"
	"github.com/omjikush09/kiranaClub/database"
	"github.com/omjikush09/kiranaClub/routes"
	"github.com/omjikush09/kiranaClub/seed"
	"github.com/robfig/cron/v3"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	//Database
	database.Connect()
	database.Migrate()

	//Seed the db
	if err := seed.SeedStores(database.DB); err != nil {
		log.Fatalf("failed to seed stores: %v", err)
	}

	c := cron.New()

	// Schedule the job to run every 5 minutes
	_, err := c.AddFunc("* * * * *", cronJob.HandleJobs)
	if err != nil {
		log.Fatalf("failed to start corn: %v", err)
	}
	// Start the cron job scheduler
	c.Start()

	//Server
	app := fiber.New()

	api := app.Group("/api")
	//Routes
	routes.RegisterJobRoutes(api)

	port := config.GetEnv("PORT", "8080")

	log.Fatal(app.Listen(":" + port))

}
