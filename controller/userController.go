package controller

import (
	"crypto/sha256"
	"fmt"
	"foodresq/dto"
	"foodresq/entities"
	errorhandler "foodresq/errorHandler"
	"foodresq/middleware"
	"foodresq/repository"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Repo repository.UserInterface
}

func NewUserController(usrepo repository.UserInterface) *UserController {
	return &UserController{
		Repo: usrepo,
	}
}

func (uc UserController) Signup() echo.HandlerFunc {

	return func(c echo.Context) error {
		newUser := dto.SignupRequest{}
		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid request body"))
		}

		passwordHash := sha256.Sum256([]byte(newUser.Password))
		passwordString := fmt.Sprintf("%x", passwordHash[:])
		request := entities.User{
			Email:       newUser.Email,
			Password:    passwordString,
			Name:        newUser.Name,
			PhoneNumber: newUser.PhoneNumber,
			Address:     newUser.Address,
		}

		if res, err := uc.Repo.Signup(request); err != nil || res.ID == 0 {
			return c.JSON(http.StatusInternalServerError, errorhandler.ResponseWriter(err.Error()))
		} else {
			response := dto.UserResponse{
				ID:          res.ID,
				Name:        res.Name,
				Email:       res.Email,
				PhoneNumber: res.PhoneNumber,
				Address:     res.Address,
			}

			return c.JSON(http.StatusCreated, response)
		}

	}
}

func (uc UserController) Login() echo.HandlerFunc {

	return func(c echo.Context) error {
		loginUser := dto.LoginRequest{}
		if err := c.Bind(&loginUser); err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid request body"))
		}

		passwordHash := sha256.Sum256([]byte(loginUser.Password))
		passwordString := fmt.Sprintf("%x", passwordHash[:])
		if res, err := uc.Repo.Login(loginUser.Email, passwordString); err != nil || res.ID == 0 {
			return c.JSON(http.StatusInternalServerError, errorhandler.ResponseWriter(err.Error()))
		} else {
			token, _ := middleware.GenerateTokenUser(res.ID)

			return c.JSON(http.StatusOK, dto.LoginResponse{
				Message: "Login success",
				Token:   token,
			})
		}
	}
}

func (uc UserController) GetUser() echo.HandlerFunc {

	return func(c echo.Context) error {

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := uint(claims["userid"].(float64))

		if res, err := uc.Repo.Get(userID); err != nil || res.ID == 0 {
			return c.JSON(http.StatusInternalServerError, errorhandler.ResponseWriter(err.Error()))
		} else {
			response := dto.UserResponse{
				ID:          res.ID,
				Name:        res.Name,
				Email:       res.Email,
				PhoneNumber: res.PhoneNumber,
				Address:     res.Address,
			}

			responses := dto.GetResponse{
				Message: "Get user success",
				Data:    response,
			}

			return c.JSON(http.StatusOK, responses)
		}
	}
}

func (uc UserController) UpdateUser() echo.HandlerFunc {

	return func(c echo.Context) error {

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := uint(claims["userid"].(float64))

		updateUserReq := dto.UpdateUserRequest{}
		if err := c.Bind(&updateUserReq); err != nil || updateUserReq.Email == "" || updateUserReq.Password == "" {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid request body"))
		}

		passwordHash := sha256.Sum256([]byte(updateUserReq.Password))
		passwordString := fmt.Sprintf("%x", passwordHash[:])

		updateUser := entities.User{
			Email:       updateUserReq.Email,
			Password:    passwordString,
			Name:        updateUserReq.Name,
			PhoneNumber: updateUserReq.PhoneNumber,
			Address:     updateUserReq.Address,
		}
		if res, err := uc.Repo.Update(uint(userID), updateUser); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, errorhandler.ResponseWriter(err.Error()))
		} else {
			responses := dto.UserResponse{
				ID:          res.ID,
				Name:        res.Name,
				Email:       res.Email,
				PhoneNumber: res.PhoneNumber,
				Address:     res.Address,
			}

			response := dto.GetResponse{
				Message: "Update user success",
				Data:    responses,
			}

			return c.JSON(http.StatusOK, response)
		}
	}
}
