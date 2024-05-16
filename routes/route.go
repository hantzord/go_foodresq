package routes

import (
	"foodresq/controller"
	"foodresq/controller/products"
	"foodresq/controller/restaurants"
	"foodresq/controller/transaction"
	mid "foodresq/middleware"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc *controller.UserController, rc *restaurants.RestaurantsController, pc *products.ProductController, tc *transaction.TransactionController) {
	e.Use(middleware.Logger())

	user := e.Group("/v1/users")
	user.POST("/signup", uc.Signup())
	user.POST("/login", uc.Login())

	user.GET("/profile", uc.GetUser(), middleware.JWT([]byte(mid.GetSigningKey())))
	user.PUT("/profile", uc.UpdateUser(), middleware.JWT([]byte(mid.GetSigningKey())))

	restaurant := e.Group("/v1/restaurants")
	restaurant.POST("/signup", rc.SignupRestaurant())
	restaurant.POST("/login", rc.LoginRestaurant())

	restaurant.GET("/profile", rc.GetMyRestaurant(), middleware.JWT([]byte(mid.GetSigningKey())))
	restaurant.PUT("/profile", rc.UpdateMyRestaurant(), middleware.JWT([]byte(mid.GetSigningKey())))

	restaurant.POST("/profile/info", rc.CreateRestaurantInfo(), middleware.JWT([]byte(mid.GetSigningKey())))
	restaurant.PUT("/profile/info", rc.UpdateRestaurantInfo(), middleware.JWT([]byte(mid.GetSigningKey())))

	restaurant.GET("/list", rc.GetAllRestaurants())
	restaurant.GET("/validate/:uniqueCode", tc.ValidateUniqueCode(), middleware.JWT([]byte(mid.GetSigningKey())))

	product := e.Group("/v1/products")
	product.POST("/create", pc.CreateProduct(), middleware.JWT([]byte(mid.GetSigningKey())))
	product.GET("/discover", pc.GetAllProduct())
	product.PUT("/update", pc.UpdateProduct(), middleware.JWT([]byte(mid.GetSigningKey())))    // Update berdasarkan nama
	product.DELETE("/delete", pc.DeleteProduct(), middleware.JWT([]byte(mid.GetSigningKey()))) // Delete berdasarkan nama

	transaction := e.Group("/v1/transactions")
	transaction.POST("/purchase", tc.PurchaseProductByName(), middleware.JWT([]byte(mid.GetSigningKey())))
}
