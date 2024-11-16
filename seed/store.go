package seed

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/omjikush09/kiranaClub/models"
	"gorm.io/gorm"
)

// SeedStores seeds store data from a CSV file into the database.
func SeedStores(db *gorm.DB) error {

	file, err := os.Open("./seed/storeData.csv")
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	// Parse the CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to parse CSV: %w", err)
	}

	for _, record := range records {
		if len(record) < 3 {
			log.Printf("skipping invalid record: %v\n", record)
			continue
		}

		pincode, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("invalid pincode in record: %v, skipping\n", record)
			continue
		}

		store := models.Store{
			Pincode:   pincode,
			StoreName: record[1],
			StoreId:   record[2],
		}

		var existingStore models.Store
		if err := db.Where("store_id = ?", store.StoreId).First(&existingStore).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Insert the store if it does not exist
				if err := db.Create(&store).Error; err != nil {
					log.Printf("failed to insert store: %v, error: %v\n", store, err)
				} else {
					log.Printf("successfully inserted store: %v\n", store)
				}
			} else {
				return fmt.Errorf("failed to query store: %w", err)
			}
		}
	}
	fmt.Println("Database Seeding Done")
	return nil
}
