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
	repository.db.Save(&restaurantInfo)

	newRestaurant.RestaurantInfoID = restaurantInfo.ID

	if err := repository.db.Save(&newRestaurant).Error; err != nil {
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

func (repository *RestaurantRepositoryDB) CreateInfo(restaurantId uint, updateRestaurant entities.RestaurantInfo) (entities.RestaurantInfo, error) {
	var existingRestaurantInfo entities.RestaurantInfo

	// Cari informasi restoran yang sudah ada berdasarkan ID
	if err := repository.db.Where("id = ?", restaurantId).First(&existingRestaurantInfo).Error; err != nil {
		// Jika tidak ada informasi restoran yang ada, buat yang baru
		newRestaurantInfo := entities.RestaurantInfo{
			Name:        updateRestaurant.Name,
			Latitude:    updateRestaurant.Latitude,
			Longitude:   updateRestaurant.Longitude,
			City:        updateRestaurant.City,
			Address:     updateRestaurant.Address,
			PhoneNumber: updateRestaurant.PhoneNumber,
			Description: updateRestaurant.Description,
		}
		if err := repository.db.Create(&newRestaurantInfo).Error; err != nil {
			return entities.RestaurantInfo{}, err
		}
		return newRestaurantInfo, nil
	}

	// Jika informasi restoran sudah ada, lakukan pembaruan
	existingRestaurantInfo.Name = updateRestaurant.Name
	existingRestaurantInfo.Latitude = updateRestaurant.Latitude
	existingRestaurantInfo.Longitude = updateRestaurant.Longitude
	existingRestaurantInfo.City = updateRestaurant.City
	existingRestaurantInfo.Address = updateRestaurant.Address
	existingRestaurantInfo.PhoneNumber = updateRestaurant.PhoneNumber
	existingRestaurantInfo.Description = updateRestaurant.Description

	if err := repository.db.Save(&existingRestaurantInfo).Error; err != nil {
		return entities.RestaurantInfo{}, err
	}

	return existingRestaurantInfo, nil
}

func (repository *RestaurantRepositoryDB) UpdateInfo(restaurantID uint, updateRestaurant entities.RestaurantInfo) (entities.RestaurantInfo, error) {
	var existingRestaurantInfo entities.RestaurantInfo

	// Cari informasi restoran yang sudah ada berdasarkan ID restoran
	if err := repository.db.First(&existingRestaurantInfo, restaurantID).Error; err != nil {
		return entities.RestaurantInfo{}, err
	}

	// Lakukan pembaruan data
	existingRestaurantInfo.Name = updateRestaurant.Name
	existingRestaurantInfo.Latitude = updateRestaurant.Latitude
	existingRestaurantInfo.Longitude = updateRestaurant.Longitude
	existingRestaurantInfo.City = updateRestaurant.City
	existingRestaurantInfo.Address = updateRestaurant.Address
	existingRestaurantInfo.PhoneNumber = updateRestaurant.PhoneNumber
	existingRestaurantInfo.Description = updateRestaurant.Description

	// Simpan pembaruan ke database
	if err := repository.db.Save(&existingRestaurantInfo).Error; err != nil {
		return entities.RestaurantInfo{}, err
	}

	return existingRestaurantInfo, nil
}
