package mocks

import (
	"foodresq/entities"

	"github.com/stretchr/testify/mock"
)

type ProductUsecaseMock struct {
	mock.Mock
}

func (m *ProductUsecaseMock) CreateProduct(restaurantID uint, product entities.RestaurantProduct) (*entities.RestaurantProduct, error) {
	args := m.Called(restaurantID, product)
	return args.Get(0).(*entities.RestaurantProduct), args.Error(1)
}

// GetAllProduct digunakan untuk memalsukan metode GetAllProduct
func (m *ProductUsecaseMock) GetAllProduct() ([]entities.RestaurantProduct, error) {
	args := m.Called()
	return args.Get(0).([]entities.RestaurantProduct), args.Error(1)
}

// UpdateProduct digunakan untuk memalsukan metode UpdateProduct
func (m *ProductUsecaseMock) UpdateProduct(restaurantID uint, product entities.RestaurantProduct) (*entities.RestaurantProduct, error) {
	args := m.Called(restaurantID, product)
	return args.Get(0).(*entities.RestaurantProduct), args.Error(1)
}

// DeleteProduct digunakan untuk memalsukan metode DeleteProduct
func (m *ProductUsecaseMock) DeleteProduct(restaurantID uint, productName string) error {
	args := m.Called(restaurantID, productName)
	return args.Error(0)
}
