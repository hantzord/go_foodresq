package restaurants

import (
	"crypto/sha256"
	"fmt"
	"foodresq/configs"
	"foodresq/entities"
	"foodresq/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterRestaurant(t *testing.T) {
	config := configs.InitConfig()
	db := utils.InitDB(config)

	restaurantRepository := NewRestaurantRepository(db)

	t.Run("Register Restaurant Success", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "test@gmail.com"
		newRestaurant.Password = passwordString

		res, err := restaurantRepository.Signup(newRestaurant)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)
	})

	t.Run("Register Restaurant Failed", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "test@gmail.com"
		newRestaurant.Password = passwordString

		res, err := restaurantRepository.Signup(newRestaurant)
		assert.Equal(t, uint(0), res.ID)
		assert.Nil(t, err)
	})
}

func TestLoginRestaurant(t *testing.T) {
	config := configs.InitConfig()
	db := utils.InitDB(config)

	restaurantRepository := NewRestaurantRepository(db)
	t.Run("Register Restaurant Success", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "test@gmail.com"
		newRestaurant.Password = passwordString

		res, err := restaurantRepository.Signup(newRestaurant)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)
	})

	t.Run("Login Restaurant Success", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var loginRestaurant entities.Restaurant
		loginRestaurant.Email = "test@gmail.com"
		loginRestaurant.Password = passwordString

		res, err := restaurantRepository.Login(loginRestaurant.Email, loginRestaurant.Password)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)
	})

	t.Run("Login Restaurant Failed", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var loginRestaurant entities.Restaurant
		loginRestaurant.Email = "test@gmail.com"
		loginRestaurant.Password = passwordString

		res, err := restaurantRepository.Login(loginRestaurant.Email, loginRestaurant.Password)
		assert.Equal(t, uint(0), res.ID)
		assert.Nil(t, err)
	})

}

func TestGetRestaurant(t *testing.T) {
	config := configs.InitConfig()
	db := utils.InitDB(config)

	restaurantRepository := NewRestaurantRepository(db)
	t.Run("Register Restaurant Success", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "test@gmail.com"
		newRestaurant.Password = passwordString

		res, err := restaurantRepository.Signup(newRestaurant)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)
	})

	t.Run("Create Restaurant Info Success", func(t *testing.T) {
		var updatedRestaurant entities.RestaurantInfo
		updatedRestaurant.Name = "Restaurant Name"
		updatedRestaurant.Latitude = 123.456
		updatedRestaurant.Longitude = 123.456
		updatedRestaurant.City = "Jakarta"
		updatedRestaurant.Address = "Jl. Raya"
		updatedRestaurant.PhoneNumber = "08123456789"
		updatedRestaurant.Description = "Restaurant Description"
		res, err := restaurantRepository.CreateInfo(uint(1), updatedRestaurant)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)
	})

	t.Run("Get Restaurant Info Success", func(t *testing.T) {
		res, rest, err := restaurantRepository.Get(uint(1))
		assert.Equal(t, uint(1), res.ID)
		assert.Equal(t, "Restaurant Name", rest.Name)
		assert.Nil(t, err)
	})
}

func TestUpdateRestaurant(t *testing.T) {
	config := configs.InitConfig()
	db := utils.InitDB(config)

	restaurantRepository := NewRestaurantRepository(db)

	t.Run("Register Restaurant Success", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "test@gmail.com"
		newRestaurant.Password = passwordString

		res, err := restaurantRepository.Signup(newRestaurant)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)
	})

	t.Run("Update Restaurant Success", func(t *testing.T) {
		var updatedRestaurant entities.Restaurant
		updatedRestaurant.Email = "test@gmail.com"
		updatedRestaurant.Password = "password"
		res, err := restaurantRepository.Update(uint(1), updatedRestaurant)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)
	})

	t.Run("Update Restaurant Failed", func(t *testing.T) {
		var updatedRestaurant entities.Restaurant
		updatedRestaurant.Email = "test@gmail.com"
		updatedRestaurant.Password = "password"
		res, err := restaurantRepository.Update(uint(0), updatedRestaurant)
		assert.Equal(t, uint(0), res.ID)
		assert.Nil(t, err)
	})

}

func TestUpdateRestaurantInfo(t *testing.T) {
	config := configs.InitConfig()
	db := utils.InitDB(config)

	restaurantRepository := NewRestaurantRepository(db)
	t.Run("Register Restaurant Success", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "test@gmail.com"
		newRestaurant.Password = passwordString

		res, err := restaurantRepository.Signup(newRestaurant)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)
	})

	t.Run("Update Restaurant Info Success", func(t *testing.T) {
		var updatedRestaurant entities.RestaurantInfo
		updatedRestaurant.Name = "Restaurant Name"
		updatedRestaurant.Latitude = 123.456
		updatedRestaurant.Longitude = 123.456
		updatedRestaurant.City = "Jakarta"
		updatedRestaurant.Address = "Jl. Raya"
		updatedRestaurant.PhoneNumber = "08123456789"
		updatedRestaurant.Description = "Restaurant Description"
		res, err := restaurantRepository.UpdateInfo(uint(1), updatedRestaurant)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)
	})

	t.Run("Update Restaurant Info Failed", func(t *testing.T) {
		var updatedRestaurant entities.RestaurantInfo
		updatedRestaurant.Name = "Restaurant Name"
		updatedRestaurant.Latitude = 123.456
		updatedRestaurant.Longitude = 123.456
		updatedRestaurant.City = "Jakarta"
		updatedRestaurant.Address = "Jl. Raya"
		updatedRestaurant.PhoneNumber = "08123456789"
		updatedRestaurant.Description = "Restaurant Description"
		res, err := restaurantRepository.UpdateInfo(uint(0), updatedRestaurant)
		assert.Equal(t, uint(0), res.ID)
		assert.Equal(t, "record not found", err.Error())
	})
}

func TestRestaurantInfo(t *testing.T) {
	config := configs.InitConfig()
	db := utils.InitDB(config)

	restaurantRepository := NewRestaurantRepository(db)

	t.Run("Register Restaurant Success", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "test@gmail.com"
		newRestaurant.Password = passwordString

		res, err := restaurantRepository.Signup(newRestaurant)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)
	})

	t.Run("Create Restaurant Info Success", func(t *testing.T) {
		var newRestaurant entities.RestaurantInfo
		newRestaurant.Name = "Restaurant Name"
		newRestaurant.Latitude = 123.456
		newRestaurant.Longitude = 123.456
		newRestaurant.City = "Bekasi"
		newRestaurant.Address = "Jl. Raya"
		newRestaurant.PhoneNumber = "08123456789"
		newRestaurant.Description = "Restaurant Description"
		res, err := restaurantRepository.CreateInfo(uint(1), newRestaurant)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)
	})

	t.Run("Create Restaurant Info Failed", func(t *testing.T) {
		var newRestaurant entities.RestaurantInfo
		newRestaurant.Name = "Restaurant Name"
		newRestaurant.Latitude = 123.456
		newRestaurant.Longitude = 123.456
		newRestaurant.City = "Bekasi"
		newRestaurant.Address = "Jl. Raya"
		newRestaurant.PhoneNumber = "08123456789"
		newRestaurant.Description = "Restaurant Description"
		res, err := restaurantRepository.CreateInfo(uint(2), newRestaurant)
		assert.Equal(t, uint(0), res.ID)
		assert.Equal(t, "record not found", err.Error())
	})
}

func TestGetAllRestaurant(t *testing.T) {
	config := configs.InitConfig()
	db := utils.InitDB(config)

	restaurantRepository := NewRestaurantRepository(db)

	t.Run("Get All Restaurant Success", func(t *testing.T) {
		restaurants, err := restaurantRepository.GetAll()
		assert.Nil(t, err)
		assert.NotNil(t, restaurants)
	})
}
