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

func (pr *ProductRepositoryDB) GetAllProduct() ([]entities.RestaurantProduct, error) {
	var products []entities.RestaurantProduct
	if err := pr.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *ProductRepositoryDB) UpdateProduct(restaurantID uint, updatedProduct entities.RestaurantProduct) (entities.RestaurantProduct, error) {
	// Query untuk mencari produk yang akan diperbarui
	var existingProduct entities.RestaurantProduct
	if err := pr.db.Where("restaurant_info_id = ? AND product_name = ?", restaurantID, updatedProduct.ProductName).First(&existingProduct).Error; err != nil {
		return entities.RestaurantProduct{}, err
	}

	// Update atribut produk yang ada
	existingProduct.Quantity = updatedProduct.Quantity
	existingProduct.ProductCategory = updatedProduct.ProductCategory
	existingProduct.ProductPrice = updatedProduct.ProductPrice
	existingProduct.ProductCondition = updatedProduct.ProductCondition
	existingProduct.ProductImage = updatedProduct.ProductImage
	existingProduct.ExpiryDate = updatedProduct.ExpiryDate
	existingProduct.Description = updatedProduct.Description

	// Simpan perubahan ke dalam database
	if err := pr.db.Save(&existingProduct).Error; err != nil {
		return entities.RestaurantProduct{}, err
	}

	return existingProduct, nil
}

func (pr *ProductRepositoryDB) DeleteProduct(restaurantID uint, productName string) error {
	// Cek apakah produk ada dalam database
	var existingProduct entities.RestaurantProduct
	if err := pr.db.Where("restaurant_info_id = ? AND product_name = ?", restaurantID, productName).First(&existingProduct).Error; err != nil {
		return err
	}

	// Hapus produk dari database
	if err := pr.db.Delete(&existingProduct).Error; err != nil {
		return err
	}

	return nil
}
