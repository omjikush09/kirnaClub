package database

import (
	"log"

	"github.com/omjikush09/kiranaClub/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.Job{}, &models.Image{}, &models.Store{}, &models.Error{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully.")
}
