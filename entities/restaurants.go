package entities

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	ID               uint
	Email            string         `gorm:"not null;unique"`
	Password         string         `gorm:"not null"`
	RestaurantInfoID uint           `gorm:"foreignKey:RestaurantInfoID"`
	RestaurantInfo   RestaurantInfo `gorm:"constraint:OnDelete:CASCADE"`
}
