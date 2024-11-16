package cronJob

import (
	"errors"
	"fmt"
	"image/jpeg"
	"log"
	"net/http"

	"github.com/omjikush09/kiranaClub/database"
	"github.com/omjikush09/kiranaClub/models"
)

func HandleJobs() {
	// Step 2: Find all ongoing jobs
	fmt.Println("Corn is runnig to handle job...")
	var jobs []models.Job
	if err := database.DB.Where("status = ?", "ongoing").Preload("Images").Preload("Stores").Find(&jobs).Error; err != nil {
		log.Printf("Error fetching ongoing jobs: %v", err)
		return
	}

	// Step 3: Loop through each ongoing job
	for _, job := range jobs {
		processJob(&job)
	}
}

func processJob(job *models.Job) {
	// For each ongoing job, loop through related stores and images
	for _, image := range job.Images {

		// Calculate the image parameter
		perimeter, err := calculateParameter(image)

		if err != nil {
			addErrorToJob(job, &image, err.Error())
			continue
		}
		// Update the image with the calculated parameter
		image.Perimeter = perimeter
		if err := database.DB.Save(&image).Error; err != nil {
			//Cron will retry
			log.Printf("Error saving image: %v", err)
			continue

		}
	}

	//Update job status to completed if all images are processed
	if job.Status == "ongoing" {
		updateJobStatusToCompleted(job)
	}
}

func calculateParameter(image models.Image) (int, error) {

	resp, err := http.Get(image.URL)
	if err != nil {
		return 0, errors.New("failed to download image")
	}
	defer resp.Body.Close()

	// Ensure the response body is closed even if an error occurs during decoding
	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	// Decode the JPEG image
	img, err := jpeg.Decode(resp.Body)
	if err != nil {
		return 0, errors.New("failed to decode JPEG image")
	}

	// Calculate the perimeter of the image
	width := img.Bounds().Dx()  // Get the width of the image
	height := img.Bounds().Dy() // Get the height of the image

	// Perimeter of a rectangle: 2 * (width + height)
	perimeter := 2 * (width + height)

	return perimeter, nil
}

func addErrorToJob(job *models.Job, image *models.Image, message string) {
	//  Add an error record for the job
	errorRecord := models.Error{
		JobId:    job.ID,
		StoreId:  image.Store.StoreId,
		Messsage: message,
	}

	if err := database.DB.Create(&errorRecord).Error; err != nil {
		log.Printf("Error creating error record: %v", err)
	}
}

func updateJobStatusToCompleted(job *models.Job) {
	//  Update job status to "completed"
	job.Status = "completed"
	if err := database.DB.Save(job).Error; err != nil {
		log.Printf("Error updating job status to completed: %v", err)
	}
}
