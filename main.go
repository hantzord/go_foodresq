package main

import (
	"fmt"
	"foodresq/configs"
	"foodresq/controller"
	"foodresq/controller/products"
	"foodresq/controller/restaurants"
	"foodresq/repository"
	productRepository "foodresq/repository/products"
	restaurantRepository "foodresq/repository/restaurants"
	"foodresq/routes"
	"foodresq/utils"

	"github.com/labstack/echo/v4"
)

func main() {

	config := configs.InitConfig()
	db := utils.InitDB(config)

	e := echo.New()
	userRepository := repository.NewUserRepository(db)
	UserController := controller.NewUserController(userRepository)

	restaurantRepository := restaurantRepository.NewRestaurantRepository(db)
	restaurantsController := restaurants.NewRestaurantController(restaurantRepository)

	productRepository := productRepository.NewProductRepository(db)
	productController := products.NewProductController(productRepository)

	routes.InitRoute(e, UserController, restaurantsController, productController)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
