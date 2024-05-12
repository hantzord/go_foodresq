package products

import (
	"foodresq/entities"

	"gorm.io/gorm"
)

type ProductRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryDB {
	return &ProductRepositoryDB{
		db: db,
	}
}

// CreateProduct method to create new product
func (pr *ProductRepositoryDB) CreateProduct(restaurantID uint, newProduct entities.RestaurantProduct) (entities.RestaurantProduct, error) {
	// Set RestaurantInfoID untuk produk yang baru
	newProduct.RestaurantInfoID = restaurantID

	// Simpan produk ke dalam database
	if err := pr.db.Create(&newProduct).Error; err != nil {
		return entities.RestaurantProduct{}, err
	}

	return newProduct, nil
}
