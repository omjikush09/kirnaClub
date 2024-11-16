package models

import "gorm.io/gorm"

type Error struct {
	gorm.Model
	JobId    uint   `gorm:"not null"`
	Job      Job    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StoreId  string `json:"store_id"`
	Messsage string `json:"message"`
}
