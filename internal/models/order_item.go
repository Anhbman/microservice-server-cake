package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID    uint64 `json:"order_id" gorm:"not null;uniqueIndex:order_cake_idx"`
	CakeID     uint64 `json:"cake_id" gorm:"not null;uniqueIndex:order_cake_idx"`
	Quantity   int64  `json:"quantity" gorm:"not null"`
	Price      int64  `json:"price" gorm:"not null"`
	Decription string `json:"description"`
	Cake       Cake   `json:"cake"`
	Order      Order  `json:"order"`
}
