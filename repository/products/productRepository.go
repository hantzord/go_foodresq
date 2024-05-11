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
func (pr *ProductRepositoryDB) CreateProduct(newProduct entities.RestaurantProduct) (entities.RestaurantProduct, error) {
	product := entities.RestaurantProduct{}

	if err := pr.db.Where("restaurant_info_id = ?", newProduct.RestaurantInfoID).Create(&newProduct).Error; err != nil {
		return product, err
	} else {
		return newProduct, nil
	}
}
