package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	ImageURL    string `json:"image_url"`
}
