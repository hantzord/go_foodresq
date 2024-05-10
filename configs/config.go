package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port        string
	DatabaseURL struct {
		DBusername string
		DBpassword string
		DBhost     string
		DBport     string
		DBname     string
	}
}

func InitConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg AppConfig
	cfg.Port = os.Getenv("APP_PORT")
	cfg.DatabaseURL.DBusername = os.Getenv("DB_USERNAME")
	cfg.DatabaseURL.DBpassword = os.Getenv("DB_PASSWORD")
	cfg.DatabaseURL.DBhost = os.Getenv("DB_HOST")
	cfg.DatabaseURL.DBport = os.Getenv("DB_PORT")
	cfg.DatabaseURL.DBname = os.Getenv("DB_NAME")

	return &cfg
}
