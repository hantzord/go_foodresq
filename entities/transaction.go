package entities

import (
	"time"

	"gorm.io/gorm"
)

// Transaction represents a payment transaction for product purchases.
type Transaction struct {
	gorm.Model
	ID                  uint
	UserID              uint
	RestaurantProductID uint
	TotalAmount         float64
	UniqueCode          string
	PaymentType         string
	PickupTime          time.Time
	CreatedAt           time.Time
}

func NewTransaction(userID uint, restaurantProductID uint, totalAmount float64, uniqueCode string, paymentType string, pickupTime time.Time) *Transaction {
	return &Transaction{
		UserID:              userID,
		RestaurantProductID: restaurantProductID,
		TotalAmount:         totalAmount,
		UniqueCode:          uniqueCode,
		PaymentType:         paymentType,
		PickupTime:          pickupTime,
	}
}
