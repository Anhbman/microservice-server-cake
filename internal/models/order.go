package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CakeID   int64 `json:"cake_id" gorm:"not null"`
	Cake     Cake  `json:"cake"`
	UserID   int64 `json:"user_id" gorm:"not null"`
	User     User  `json:"user"`
	Quantity int64 `json:"quantity" gorm:"not null"`
	Price    int64 `json:"price" gorm:"not null"`
}
