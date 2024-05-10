package main

import (
	"fmt"
	"foodresq/configs"
	"foodresq/controller"
	"foodresq/controller/restaurants"
	"foodresq/repository"
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

	routes.InitRoute(e, UserController, restaurantsController)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
