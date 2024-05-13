package mocks

import (
	"foodresq/entities"

	"github.com/stretchr/testify/mock"
)

type RestaurantUsecaseMock struct {
	mock.Mock
}

func (m *RestaurantUsecaseMock) Signup(restaurant entities.Restaurant) (*entities.Restaurant, error) {
	args := m.Called(restaurant)
	return args.Get(0).(*entities.Restaurant), args.Error(1)
}

// Login digunakan untuk memalsukan metode Login
func (m *RestaurantUsecaseMock) Login(email, password string) (*entities.Restaurant, error) {
	args := m.Called(email, password)
	return args.Get(0).(*entities.Restaurant), args.Error(1)
}

// Update digunakan untuk memalsukan metode Update
func (m *RestaurantUsecaseMock) Update(id uint, restaurant entities.Restaurant) (*entities.Restaurant, error) {
	args := m.Called(id, restaurant)
	return args.Get(0).(*entities.Restaurant), args.Error(1)
}

// Get digunakan untuk memalsukan metode Get
func (m *RestaurantUsecaseMock) Get(id uint) (*entities.Restaurant, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Restaurant), args.Error(1)
}

// CreateInfo digunakan untuk memalsukan metode CreateInfo
func (m *RestaurantUsecaseMock) CreateInfo(restaurantID uint, info entities.RestaurantInfo) (*entities.RestaurantInfo, error) {
	args := m.Called(restaurantID, info)
	return args.Get(0).(*entities.RestaurantInfo), args.Error(1)
}

// UpdateInfo digunakan untuk memalsukan metode UpdateInfo
func (m *RestaurantUsecaseMock) UpdateInfo(restaurantID uint, info entities.RestaurantInfo) (*entities.RestaurantInfo, error) {
	args := m.Called(restaurantID, info)
	return args.Get(0).(*entities.RestaurantInfo), args.Error(1)
}

// GetAll digunakan untuk memalsukan metode GetAll
func (m *RestaurantUsecaseMock) GetAll() ([]entities.RestaurantInfo, error) {
	args := m.Called()
	return args.Get(0).([]entities.RestaurantInfo), args.Error(1)
}
