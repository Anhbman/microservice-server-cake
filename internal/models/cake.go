package models

import (
	"gorm.io/gorm"
)

type Cake struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       int64 `json:"price"`
	ImageUrl	  string  `json:"image_url"`
	UserID      int64     `json:"user_id"`
	User        User    `json:"user"`
}

