package middleware

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GetSigningKey() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}

func GenerateTokenUser(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = id
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(GetSigningKey())
}

func GenerateTokenRestaurant(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["restaurantid"] = id
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(GetSigningKey())
}
