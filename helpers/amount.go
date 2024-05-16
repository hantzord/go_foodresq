package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"foodresq/entities"
)

func CalculateTotalAmount(products []entities.RestaurantProduct) float64 {
	total := 0.0
	for _, p := range products {
		total += float64(p.ProductPrice)
	}
	return total
}

func GenerateUniqueCode() string {
	bytes := make([]byte, 4)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
