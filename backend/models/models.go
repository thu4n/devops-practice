package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	Email  string
	Orders []Order
}

type Drink struct {
	gorm.Model
	Name        string
	Description string
	DrinkSizes  []DrinkSize // Each drink can have multiple sizes
}

type DrinkSize struct {
	gorm.Model
	DrinkID uint    `gorm:"foreignKey:DrinkID;references:ID"`
	Size    string  // "Small", "Medium", "Large"
	Price   float32 // Price for this size
}

type Order struct {
	gorm.Model
	UserID      uint `gorm:"foreignKey:UserID;references:ID"`
	OrderDate   time.Time
	OrderStatus string
	OrderItems  []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID     uint `gorm:"foreignKey:OrderID;references:ID"`
	DrinkSizeID uint `gorm:"foreignKey:DrinkSizeID;references:ID"`
	Quantity    uint
	Price       float32
}
