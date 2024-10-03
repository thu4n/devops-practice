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
	// No password needed for this simple app
}

type Drink struct {
	gorm.Model
	Name        string
	Description string
	Size        string
	Price       float32
}

type Order struct {
	gorm.Model
	UserID      uint `gorm:"foreignKey:UserID;references:ID"`
	OrderDate   time.Time
	Size        string
	OrderStatus string
	OrderItems  []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID  uint `gorm:"foreignKey:OrderID;references:ID"`
	DrinkID  uint `gorm:"foreignKey:DrinkID;references:ID"`
	Quantity uint
	Price    float32
}
