package entities

import (
	"fmt"
	"strings"
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

type RestaurantProduct struct {
	gorm.Model       `json:"-"`
	ID               uint
	RestaurantInfoID uint
	ProductName      string           `gorm:"not null"`
	Quantity         int              `gorm:"column:quantity;type:int;not null;default:0"`
	ProductCategory  ProductCategory  `gorm:"type:enum('Food','Drink','Snack','Dessert')"`
	ProductPrice     int              `gorm:"column:product_price;type:int;not null;default:0"`
	ProductCondition ProductCondition `gorm:"type:enum('Fresh','Stale','Rotten')"`
	ProductImage     string           `gorm:"not null"`
	ExpiryDate       time.Time        `gorm:"not null"`
	Description      string
}

// ParseProductCategory mengonversi string menjadi ProductCategory enum
func ParseProductCategory(category string) (ProductCategory, error) {
	switch strings.ToLower(category) {
	case "food":
		return Food, nil
	case "drink":
		return Drink, nil
	case "snack":
		return Snack, nil
	case "dessert":
		return Dessert, nil
	default:
		return "", fmt.Errorf("invalid product category: %s", category)
	}
}

// ParseProductCondition mengonversi string menjadi ProductCondition enum
func ParseProductCondition(condition string) (ProductCondition, error) {
	switch strings.ToLower(condition) {
	case "fresh":
		return Fresh, nil
	case "stale":
		return Stale, nil
	case "rotten":
		return Rotten, nil
	default:
		return "", fmt.Errorf("invalid product condition: %s", condition)
	}
}

// String mengembalikan representasi string dari ProductCategory enum
func (pc ProductCategory) String() string {
	return string(pc)
}

// String mengembalikan representasi string dari ProductCondition enum
func (pc ProductCondition) String() string {
	return string(pc)
}
