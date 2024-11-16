package database

import (
	"fmt"
	"github.com/omjikush09/kiranaClub/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	

	// Build the connection string from environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.GetEnv("DATABASE_HOST", "localhost"),
		config.GetEnv("DATABASE_USER", "postgres"),
		config.GetEnv("DATABASE_PASSWORD", "postgres"),
		config.GetEnv("DATABASE_NAME", "postgres"),
		config.GetEnv("DATABASE_PORT", "5432"),
		config.GetEnv("DATABASE_SSLMODE", "disable"),
		config.GetEnv("TIME_ZONE", "UTC"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established successfully.")
}



