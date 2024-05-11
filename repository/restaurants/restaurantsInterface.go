package restaurants

import "foodresq/entities"

type RestaurantInterface interface {
	Signup(newRestaurant entities.Restaurant) (entities.Restaurant, error)
	Login(email, password string) (entities.Restaurant, error)
	Get(restaurantId uint) (entities.Restaurant, entities.RestaurantInfo, error)
	CreateInfo(restaurantId uint, updateRestaurant entities.RestaurantInfo) (entities.RestaurantInfo, error)
	UpdateInfo(restaurantId uint, updateRestaurant entities.RestaurantInfo) (entities.RestaurantInfo, error)
	Update(restaurantId uint, updatedRestaurant entities.Restaurant) (entities.Restaurant, error)
}
