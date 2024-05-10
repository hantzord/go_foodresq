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
