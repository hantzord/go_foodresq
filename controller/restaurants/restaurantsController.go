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

func (rc RestaurantsController) UpdateMyRestaurant() echo.HandlerFunc {

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
				Message: "Email or password",
				Data:    data,
			}

			return c.JSON(http.StatusOK, response)
		}
	}
}

func (rc RestaurantsController) GetMyRestaurant() echo.HandlerFunc {

	return func(c echo.Context) error {

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restaurantId := claims["restaurantid"].(float64)

		if res, _, err := rc.Repo.Get(uint(restaurantId)); err != nil {
			return c.JSON(http.StatusNotFound, errorhandler.ResponseWriter(err.Error()))

		} else {

			restaurantInfo := dto.RestaurantInfoResponse{
				ID:                 res.RestaurantInfo.ID,
				Name:               res.RestaurantInfo.Name,
				Latitude:           res.RestaurantInfo.Latitude,
				Longitude:          res.RestaurantInfo.Longitude,
				City:               res.RestaurantInfo.City,
				Address:            res.RestaurantInfo.Address,
				PhoneNumber:        res.RestaurantInfo.PhoneNumber,
				Description:        res.RestaurantInfo.Description,
				RestaurantProducts: res.RestaurantInfo.RestaurantProducts,
			}

			response := dto.RestaurantProfileResponseMsg{
				Message: "Get restaurant profile success",
				Data:    restaurantInfo,
			}

			return c.JSON(http.StatusOK, response)
		}
	}
}

func (rc RestaurantsController) CreateRestaurantInfo() echo.HandlerFunc {

	return func(c echo.Context) error {

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restaurantId := claims["restaurantid"].(float64)

		restIReq := dto.RestaurantInfoRequest{}
		if err := c.Bind(&restIReq); err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid request body"))
		}

		restI := entities.RestaurantInfo{
			Name:        restIReq.Name,
			Latitude:    restIReq.Latitude,
			Longitude:   restIReq.Longitude,
			City:        restIReq.City,
			Address:     restIReq.Address,
			PhoneNumber: restIReq.PhoneNumber,
			Description: restIReq.Description,
		}

		if res, err := rc.Repo.CreateInfo(uint(restaurantId), restI); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, errorhandler.ResponseWriter(err.Error()))
		} else {

			restaurantInfo := dto.RestaurantInfoResponse{
				ID:          res.ID,
				Name:        res.Name,
				Latitude:    res.Latitude,
				Longitude:   res.Longitude,
				City:        res.City,
				Address:     res.Address,
				PhoneNumber: res.PhoneNumber,
				Description: res.Description,
			}

			response := dto.RestaurantInfoResponseMsg{
				Message: "Create restaurant info success",
				Data:    restaurantInfo,
			}

			return c.JSON(http.StatusOK, response)
		}
	}
}

func (rc RestaurantsController) UpdateRestaurantInfo() echo.HandlerFunc {

	return func(c echo.Context) error {

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restaurantId := claims["restaurantid"].(float64)

		updRestIReq := dto.RestaurantInfoUpdateRequest{}
		if err := c.Bind(&updRestIReq); err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid request body"))
		}

		updRest := entities.RestaurantInfo{
			Name:        updRestIReq.Name,
			Latitude:    updRestIReq.Latitude,
			Longitude:   updRestIReq.Longitude,
			City:        updRestIReq.City,
			Address:     updRestIReq.Address,
			PhoneNumber: updRestIReq.PhoneNumber,
			Description: updRestIReq.Description,
		}

		if res, err := rc.Repo.UpdateInfo(uint(restaurantId), updRest); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, errorhandler.ResponseWriter(err.Error()))
		} else {

			restaurantInfo := dto.RestaurantInfoResponse{
				ID:          res.ID,
				Name:        res.Name,
				Latitude:    res.Latitude,
				Longitude:   res.Longitude,
				City:        res.City,
				Address:     res.Address,
				PhoneNumber: res.PhoneNumber,
				Description: res.Description,
			}

			response := dto.RestaurantInfoResponseMsg{
				Message: "Update restaurant info success",
				Data:    restaurantInfo,
			}

			return c.JSON(http.StatusOK, response)
		}
	}
}

func (rc RestaurantsController) GetAllRestaurants() echo.HandlerFunc {

	return func(c echo.Context) error {
		res, err := rc.Repo.GetAll()
		if err != nil {
			return c.JSON(http.StatusNotFound, errorhandler.ResponseWriter(err.Error()))
		}

		var resp []dto.AllRestaurantListResponse

		for _, restI := range res {
			resp = append(resp, dto.AllRestaurantListResponse{
				ID:                 restI.ID,
				Name:               restI.Name,
				Latitude:           restI.Latitude,
				Longitude:          restI.Longitude,
				City:               restI.City,
				Address:            restI.Address,
				PhoneNumber:        restI.PhoneNumber,
				Description:        restI.Description,
				RestaurantProducts: restI.RestaurantProducts,
			})
		}
		response := dto.AllRestaurantListResponseMsg{
			Message: "Get all restaurants success",
			Data:    resp,
		}

		return c.JSON(http.StatusOK, response)
	}

}
