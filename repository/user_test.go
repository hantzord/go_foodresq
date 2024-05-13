package repository

import (
	"crypto/sha256"
	"fmt"
	"foodresq/configs"
	"foodresq/entities"
	"foodresq/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	config := configs.InitConfig()
	db := utils.InitDB(config)

	userRepository := NewUserRepository(db)

	t.Run("Register User Success", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.Email = "test@gmail.com"
		newUser.Password = passwordString
		newUser.Name = "test"
		newUser.PhoneNumber = "08123456789"
		newUser.Address = "test"

		res, err := userRepository.Signup(newUser)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)

	})

	t.Run("Register User Failed", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.ID = 1
		newUser.Email = "test@gmail.com"
		newUser.Password = passwordString
		newUser.Name = "test"
		newUser.PhoneNumber = "08123456789"
		newUser.Address = "test"

		res, err := userRepository.Signup(newUser)
		assert.Equal(t, uint(1), res.ID)
		assert.Error(t, err)
	})
}

func TestLoginUser(t *testing.T) {
	config := configs.InitConfig()
	db := utils.InitDB(config)

	userRepository := NewUserRepository(db)

	t.Run("Reguster User Success", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.Email = "test@gmail.com"
		newUser.Password = passwordString
		newUser.Name = "test"
		newUser.PhoneNumber = "08123456789"
		newUser.Address = "test"

		res, err := userRepository.Signup(newUser)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)
	})

	t.Run("Login User Success", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var loginuser entities.User
		loginuser.Email = "test@gmail.com"
		loginuser.Password = passwordString

		res, err := userRepository.Login(loginuser.Email, loginuser.Password)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)

		t.Run("Login User Failed", func(t *testing.T) {
			hash := sha256.Sum256([]byte("password"))
			passwordString := fmt.Sprintf("%x", hash[:])
			var loginuser entities.User
			loginuser.Email = "test@gmail.com"
			loginuser.Password = passwordString

			res, err := userRepository.Login(loginuser.Email, loginuser.Password)
			assert.Equal(t, uint(0), res.ID)
			assert.Error(t, err)
		})

	})
}

func TestGetUser(t *testing.T) {
	config := configs.InitConfig()
	db := utils.InitDB(config)

	userRepository := NewUserRepository(db)

	t.Run("Register User Success", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.Email = "test@gmail.com"
		newUser.Password = passwordString
		newUser.Name = "test"
		newUser.PhoneNumber = "08123456789"
		newUser.Address = "test"

		res, err := userRepository.Signup(newUser)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)

	})

	t.Run("Get User Success", func(t *testing.T) {
		res, err := userRepository.Get(1)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)
	})

	t.Run("Get User Failed", func(t *testing.T) {
		res, err := userRepository.Get(0)
		assert.Equal(t, uint(0), res.ID)
		assert.Error(t, err)
	})

}

func TestUpdateUser(t *testing.T) {
	config := configs.InitConfig()
	db := utils.InitDB(config)

	userRepository := NewUserRepository(db)

	t.Run("Register User Success", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.Email = "test@gmail.com"
		newUser.Password = passwordString
		newUser.Name = "test"
		newUser.PhoneNumber = "08123456789"
		newUser.Address = "test"

		res, err := userRepository.Signup(newUser)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)

	})

	t.Run("Update User Success", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var updateUser entities.User
		updateUser.Email = "test@gmail.com"
		updateUser.Password = passwordString
		updateUser.Name = "keren"
		updateUser.PhoneNumber = "08123456789"
		updateUser.Address = "jakarta"

		res, err := userRepository.Update(uint(1), updateUser)
		assert.Equal(t, uint(1), res.ID)
		assert.Nil(t, err)
	})

	t.Run("Update User Failed", func(t *testing.T) {
		hash := sha256.Sum256([]byte("password"))
		passwordString := fmt.Sprintf("%x", hash[:])
		var updateUser entities.User
		updateUser.Email = "test@gmail.com"
		updateUser.Password = passwordString
		updateUser.Name = "keren"
		updateUser.PhoneNumber = "08123456789"
		updateUser.Address = "jakarta"

		res, err := userRepository.Update(uint(5), updateUser)
		assert.Equal(t, uint(0), res.ID)
		assert.Error(t, err)
	})
}
