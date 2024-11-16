package models

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	URL       string `json:"image_url"`
	Perimeter int
	JobId     uint  `gorm:"not null"`
	Job       Job   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StoreId   uint  `gorm:"not null"`
	Store     Store `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
