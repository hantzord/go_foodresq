package utils

import (
	"fmt"
	"foodresq/configs"
	"foodresq/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DatabaseURL.DBusername,
		config.DatabaseURL.DBpassword,
		config.DatabaseURL.DBhost,
		config.DatabaseURL.DBport,
		config.DatabaseURL.DBname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	InitMigration(db)
	return db
}

func InitMigration(db *gorm.DB) {

	db.AutoMigrate(entities.User{})
	db.AutoMigrate(entities.Restaurant{})
	db.AutoMigrate(entities.RestaurantInfo{})
	db.AutoMigrate(entities.RestaurantProducts{})
	db.AutoMigrate(entities.Transaction{})

}
