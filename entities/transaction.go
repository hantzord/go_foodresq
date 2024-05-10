package entities

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID           uint
	UserID       uint `gorm:"not null"`
	RestaurantID uint `gorm:"not null"`
	DateTime     time.Time
	Persons      int    `gorm:"not null;default:1"`
	Total        int    `gorm:"not null"`
	Status       string `gorm:"NOT NULL;default:waiting for confirmation"`
	User         User
	Restaurant   Restaurant
}
