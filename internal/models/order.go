package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID   int64 `json:"user_id" gorm:"not null"`
	User     User  `json:"user"`
}
