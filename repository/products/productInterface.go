package products

import "foodresq/entities"

type ProductInterface interface {
	CreateProduct(newProduct entities.RestaurantProduct) (entities.RestaurantProduct, error)
}
