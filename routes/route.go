package routes

import (
	"foodresq/controller"
	"foodresq/controller/restaurants"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc *controller.UserController, rc *restaurants.RestaurantsController) {
	e.Use(middleware.Logger())

	user := e.Group("/v1/users")
	user.POST("/signup", uc.Signup())
	user.POST("/login", uc.Login())

	restaurant := e.Group("/v1/restaurants")
	restaurant.POST("/signup", rc.SignupRestaurant())
	restaurant.POST("/login", rc.LoginRestaurant())
}
