package restaurants

import (
	"foodresq/entities"

	"gorm.io/gorm"
)

type RestaurantRepositoryDB struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) *RestaurantRepositoryDB {
	return &RestaurantRepositoryDB{
		db: db,
	}
}

func (repository *RestaurantRepositoryDB) Signup(newRestaurant entities.Restaurant) (entities.Restaurant, error) {
	restaurantInfo := entities.RestaurantInfo{
		Name: "Restaurant Name",
	}
	repository.db.Create(&restaurantInfo)

	newRestaurant.RestaurantInfoID = restaurantInfo.ID

	if err := repository.db.Create(&newRestaurant).Error; err != nil {
		return newRestaurant, err
	} else {
		return newRestaurant, nil
	}
}

func (repository *RestaurantRepositoryDB) Login(email, password string) (entities.Restaurant, error) {
	var restaurant entities.Restaurant

	if err := repository.db.Where("email = ? AND password = ?", email, password).First(&restaurant).Error; err != nil {
		return restaurant, err
	} else {
		return restaurant, nil
	}
}

func (repository *RestaurantRepositoryDB) Get(restaurantId uint) (entities.Restaurant, entities.RestaurantInfo, error) {
	restaurant := entities.Restaurant{}
	restaurantInfo := entities.RestaurantInfo{}

	if err := repository.db.Preload("RestaurantInfo.RestaurantProducts").First(&restaurant, restaurantId).Error; err != nil {
		return restaurant, restaurantInfo, err
	} else {
		return restaurant, restaurantInfo, nil
	}
}

func (repository *RestaurantRepositoryDB) Update(restaurantId uint, updatedRestaurant entities.Restaurant) (entities.Restaurant, error) {
	restaurant := entities.Restaurant{}

	if err := repository.db.First(&restaurant, "id = ?", restaurantId).Omit("id").Error; err != nil {
		return restaurant, err
	} else {
		repository.db.Model(&restaurant).Updates(updatedRestaurant)
		return restaurant, nil
	}
}

func (repository *RestaurantRepositoryDB) Delete(restaurantId uint) (entities.Restaurant, error) {
	restaurant := entities.Restaurant{}

	if err := repository.db.First(&restaurant, "id = ?", restaurantId).Delete(&restaurant).Error; err != nil {
		return restaurant, err
	} else {
		return restaurant, nil
	}
}

func (repository *RestaurantRepositoryDB) CreateInfo(restaurantId uint, updateRestaurant entities.RestaurantInfo) (entities.RestaurantInfo, error) {
	restaurant := entities.Restaurant{}
	restaurantInfo := entities.RestaurantInfo{}

	if err := repository.db.First(&restaurant, "id = ?", restaurantId).Omit("id").Error; err != nil {
		return restaurantInfo, err
	} else {
		repository.db.First(&restaurantInfo, "id = ?", restaurant.RestaurantInfoID)
		repository.db.Model(&restaurant).Updates(updateRestaurant)
		return updateRestaurant, nil
	}

}

func (repository *RestaurantRepositoryDB) UpdateInfo(restaurantId uint, updateRestaurant entities.RestaurantInfo) (entities.RestaurantInfo, error) {
	restaurant := entities.Restaurant{}
	restaurantInfo := entities.RestaurantInfo{}

	if err := repository.db.First(&restaurant, "id = ?", restaurantId).Omit("id").Error; err != nil {
		return restaurantInfo, err
	} else {
		repository.db.First(&restaurantInfo).Updates(updateRestaurant)
		return updateRestaurant, nil
	}

}
