package repository

import "foodresq/entities"

type UserInterface interface {
	Login(email, password string) (entities.User, error)
	Signup(newUser entities.User) (entities.User, error)
}
