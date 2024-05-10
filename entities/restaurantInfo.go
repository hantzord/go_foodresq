package entities

import "gorm.io/gorm"

type RestaurantInfo struct {
	gorm.Model
	ID                 uint
	Name               string
	Price              string
	Latitude           float64
	Longitude          float64
	City               string
	Address            string
	PhoneNumber        string
	Description        string
	Restaurant         []Restaurant
	RestaurantProducts []RestaurantProducts
}
