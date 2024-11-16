package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Status    string   `gorm:"default:ongoing"`
	Stores    []*Store `gorm:"many2many:job_store;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Errors    []*Error `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Images     []Image `gorm:"foreignKey:JobId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
