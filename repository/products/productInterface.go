package products

import "foodresq/entities"

type ProductInterface interface {
	CreateProduct(restaurantID uint, newProduct entities.RestaurantProduct) (entities.RestaurantProduct, error)
}
