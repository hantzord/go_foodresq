package entities

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategory string

const (
	Food    ProductCategory = "Food"
	Drink   ProductCategory = "Drink"
	Snack   ProductCategory = "Snack"
	Dessert ProductCategory = "Dessert"
)

type ProductCondition string

const (
	Fresh  ProductCondition = "Fresh"
	Stale  ProductCondition = "Stale"
	Rotten ProductCondition = "Rotten"
)

type RestaurantProducts struct {
	gorm.Model
	ID               uint
	RestaurantInfoID uint             `gorm:"primarykey"`
	ProductName      string           `gorm:"not null"`
	ProductCategory  ProductCategory  `gorm:"type:enum('Food','Drink','Snack','Dessert')"`
	ProductPrice     int              `gorm:"column:product_price;type:int;not null;default:0"`
	ProductCondition ProductCondition `gorm:"type:enum('Fresh','Stale','Rotten')"`
	ProductImage     string           `gorm:"not null"`
	ExpiryDate       time.Time        `gorm:"not null"`
	Description      string
	RestaurantInfo   RestaurantInfo
}
