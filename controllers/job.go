package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/omjikush09/kiranaClub/database"
	"github.com/omjikush09/kiranaClub/models"
	"github.com/omjikush09/kiranaClub/utils"
)

type visit struct {
	StoreId    string   `json:"store_id"`
	ImageUrls  []string `json:"image_url"`
	Visit_time string   `json:"visit_time"`
}

type jobBody struct {
	Count  uint    `json:"count"`
	Visits []visit `json:"visits"`
}

type ErrorResponseFormat struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}

func CreateJob(c *fiber.Ctx) error {

	// var job models.Job
	var jobBody jobBody
	if err := c.BodyParser(&jobBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
			"erro":  err,
		})
	}

	if jobBody.Count != uint(len(jobBody.Visits)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Count not matched with visits",
		})
	}
	job := models.Job{Status: "ongoing"}

	database.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&job).Error; err != nil {
			return utils.InternalServerError(c)

		}

		for _, visit := range jobBody.Visits {

			if !utils.IsValidUTC(visit.Visit_time) {
					error := models.Error{
						JobId:    job.ID,
						StoreId:  visit.StoreId,
						Messsage: fmt.Sprintf("Store with ID %s  don't have visit time in UTC", visit.StoreId),
					}

					if err := tx.Create(&error).Error; err != nil {
						return utils.InternalServerError(c)
					}

					job.Status = "failed"
					if err := tx.Save(job).Error; err != nil {
						return utils.InternalServerError(c)
					}
			}

			var store models.Store
			if err := tx.Where("store_id = ?", visit.StoreId).First(&store).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// Create an error record for the missing store

					error := models.Error{
						JobId:    job.ID,
						StoreId:  visit.StoreId,
						Messsage: fmt.Sprintf("Store with ID %s not found", visit.StoreId),
					}

					if err := tx.Create(&error).Error; err != nil {
						return utils.InternalServerError(c)
					}

					job.Status = "failed"
					if err := tx.Save(job).Error; err != nil {
						return utils.InternalServerError(c)
					}

					// Skip further processing for this store
					continue
				}
				return utils.InternalServerError(c)
			}
			store.VisitTime = visit.Visit_time
			if err := tx.Save(&store).Error; err != nil {
				return utils.InternalServerError(c)
			}
			// Associate store with the job
			if err := tx.Model(&job).Association("Stores").Append(&store); err != nil {
				return utils.InternalServerError(c)
			}
			var images []models.Image
			for _, imageUrl := range visit.ImageUrls {
				images := append(images, models.Image{
					StoreId: uint(store.ID),
					URL:     imageUrl,
					JobId:   job.ID,
				})
				if len(images) > 0 {
					if err := tx.Create(&images).Error; err != nil {
						return utils.InternalServerError(c)
					}
				}

			}

		}
		return nil
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"job_id": job.ID,
	})
}

func GetJob(c *fiber.Ctx) error {

	jobId := c.Query("jobId")

	var job models.Job

	err := database.DB.Preload("Errors").Find(&job, jobId).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}

	if job.Status == "failed" {

		var errData []ErrorResponseFormat
		for _, err := range job.Errors {
			errData = append(errData, ErrorResponseFormat{
				StoreID: err.StoreId,
				Error:   err.Messsage,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": job.Status,
			"job_id": job.ID,
			"error":  errData,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": job.Status,
		"job_id": job.ID,
	})

}

func UpdateJobStatusToFailed(job *models.Job) {
	//Update job status to "failed"
	job.Status = "failed"
	if err := database.DB.Save(job).Error; err != nil {
		log.Printf("Error updating job status to failed: %v", err)
	}
}
