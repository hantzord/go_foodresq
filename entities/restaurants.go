package entities

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	ID               uint
	Email            string `gorm:"not null;unique"`
	Password         string `gorm:"not null"`
	RestaurantInfoID uint
	RestaurantInfo   RestaurantInfo `gorm:"constraint:OnDelete:CASCADE"`
}
