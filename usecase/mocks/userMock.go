package mocks

import (
	"foodresq/entities"

	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (u *UserUsecaseMock) Signup(user entities.User) (entities.User, error) {
	args := u.Called(user)
	return args.Get(0).(entities.User), args.Error(1)
}

func (u *UserUsecaseMock) Login(email, password string) (entities.User, error) {
	args := u.Called(email, password)
	return args.Get(0).(entities.User), args.Error(1)
}

func (u *UserUsecaseMock) Get(userId uint) (entities.User, error) {
	args := u.Called(userId)
	return args.Get(0).(entities.User), args.Error(1)
}

func (u *UserUsecaseMock) Update(userId uint, updateUser entities.User) (entities.User, error) {
	args := u.Called(userId, updateUser)
	return args.Get(0).(entities.User), args.Error(1)
}
