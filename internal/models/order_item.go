package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID  uint64 `json:"order_id" gorm:"not null"`
	CakeID   uint64 `json:"cake_id" gorm:"not null"`
	Quantity int64    `json:"quantity" gorm:"not null"`
	Price    int64  `json:"price" gorm:"not null"`
	Cake     Cake   `json:"cake"`
	Order    Order  `json:"order"`
}
