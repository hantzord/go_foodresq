package helpers

import (
	"foodresq/entities"

	"github.com/google/uuid"
)

func CalculateTotalAmount(products []entities.RestaurantProduct) float64 {
	total := 0.0
	for _, p := range products {
		total += float64(p.ProductPrice)
	}
	return total
}

func GenerateUniqueCode() string {
	// Membuat UUID v4 baru
	uuid := uuid.New()

	// Mengonversi UUID menjadi string dan mengambil 8 karakter pertama
	code := uuid.String()[:8]

	return code
}
