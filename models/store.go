package models

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	StoreId   string  `gorm:"uniqueIndex" json:"store_id" validate:"required"`
	Pincode   int     `gorm:"index" json:"pincode" validate:"required"`
	StoreName string  `validate:"required,min=1"`
	Jobs      []*Job  `gorm:"many2many:job_store;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Images    []Image `gorm:"foreignKey:StoreId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VisitTime string  `json:"visit_time"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
