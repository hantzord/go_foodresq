package products

import "foodresq/entities"

type ProductInterface interface {
	CreateProduct(restaurantID uint, newProduct entities.RestaurantProduct) (entities.RestaurantProduct, error)
	GetAllProduct() ([]entities.RestaurantProduct, error)
	UpdateProduct(restaurantID uint, updatedProduct entities.RestaurantProduct) (entities.RestaurantProduct, error)
	DeleteProduct(restaurantID uint, productID uint) error
}
