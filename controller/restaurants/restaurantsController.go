package restaurants

import (
	"crypto/sha256"
	"fmt"
	"foodresq/dto"
	"foodresq/entities"
	errorhandler "foodresq/errorHandler"
	"foodresq/middleware"
	"foodresq/repository/restaurants"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RestaurantsController struct {
	Repo restaurants.RestaurantInterface
}

func NewRestaurantController(repo restaurants.RestaurantInterface) *RestaurantsController {
	return &RestaurantsController{
		Repo: repo,
	}
}

func (rc RestaurantsController) SignupRestaurant() echo.HandlerFunc {
	return func(c echo.Context) error {

		newRestReq := dto.RestSignupRequest{}

		if err := c.Bind(&newRestReq); err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid request body"))
		}

		passwordHash := sha256.Sum256([]byte(newRestReq.Password))
		passwordString := fmt.Sprintf("%x", passwordHash[:])

		newRestaurant := entities.Restaurant{
			Email:    newRestReq.Email,
			Password: passwordString,
		}

		if res, err := rc.Repo.Signup(newRestaurant); err != nil || res.ID == 0 {
			return c.JSON(http.StatusInternalServerError, errorhandler.ResponseWriter(err.Error()))

		} else {
			data := dto.RestaurantResponse{
				ID:    res.ID,
				Email: res.Email,
			}

			response := dto.SignupRestaurantsResponse{
				Message: "Signup success",
				Data:    data,
			}
			return c.JSON(http.StatusOK, response)

		}
	}
}

func (rc RestaurantsController) LoginRestaurant() echo.HandlerFunc {

	return func(c echo.Context) error {
		restLogin := dto.RestLoginRequest{}

		if err := c.Bind(&restLogin); err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid request body"))
		}

		hash := sha256.Sum256([]byte(restLogin.Password))
		passwordString := fmt.Sprintf("%x", hash[:])

		if res, err := rc.Repo.Login(restLogin.Email, passwordString); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, errorhandler.ResponseWriter(err.Error()))

		} else {

			token, _ := middleware.GenerateTokenRestaurant(res.ID)

			return c.JSON(http.StatusOK, dto.RestaurantLoginResponse{
				Message: "Login success",
				Token:   token,
			})
		}

	}
}

func (rc RestaurantsController) UpdateRestaurant() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restaurantId := claims["restaurantid"].(float64)

		restoReq := dto.RestSignupRequest{}

		if err := c.Bind(&restoReq); err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid request body"))
		}

		passwordHash := sha256.Sum256([]byte(restoReq.Password))
		passwordString := fmt.Sprintf("%x", passwordHash[:])

		restaurant := entities.Restaurant{
			Email:    restoReq.Email,
			Password: passwordString,
		}

		if res, err := rc.Repo.Update(uint(restaurantId), restaurant); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, errorhandler.ResponseWriter(err.Error()))

		} else {
			data := dto.RestaurantResponse{
				ID:    res.ID,
				Email: res.Email,
			}

			response := dto.SignupRestaurantsResponse{
				Message: "Update restaurant success",
				Data:    data,
			}

			return c.JSON(http.StatusOK, response)
		}
	}
}

func (rc RestaurantsController) RestaurantProfile() echo.HandlerFunc {

	return func(c echo.Context) error {

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restaurantId := claims["restaurantid"].(float64)

		if res, resIn, err := rc.Repo.Get(uint(restaurantId)); err != nil {
			return c.JSON(http.StatusNotFound, errorhandler.ResponseWriter(err.Error()))

		} else {

			restaurantDetail := dto.RestaurantDetailResponse{
				ID:                 res.RestaurantInfo.ID,
				Name:               res.RestaurantInfo.Name,
				Latitude:           res.RestaurantInfo.Latitude,
				Longitude:          res.RestaurantInfo.Longitude,
				City:               res.RestaurantInfo.City,
				Address:            res.RestaurantInfo.Address,
				PhoneNumber:        res.RestaurantInfo.PhoneNumber,
				Description:        res.RestaurantInfo.Description,
				RestaurantProducts: resIn,
			}

			response := dto.RestaurantProfileResponse{
				Message: "Get restaurant profile success",
				Data:    res,
				Info:    restaurantDetail,
			}

			return c.JSON(http.StatusOK, response)
		}
	}
}
