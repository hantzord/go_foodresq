package repository

import (
	"foodresq/entities"

	"gorm.io/gorm"
)

type UserRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryDB {
	return &UserRepositoryDB{
		db: db,
	}
}

func (repository *UserRepositoryDB) Signup(newUser entities.User) (entities.User, error) {
	if err := repository.db.Create(&newUser).Error; err != nil {
		return newUser, err
	} else {
		return newUser, nil
	}
}

func (repository *UserRepositoryDB) Login(email, password string) (entities.User, error) {
	var user entities.User

	if err := repository.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil

}

func (repository *UserRepositoryDB) Get(userId uint) (entities.User, error) {
	user := entities.User{}

	if err := repository.db.First(&user, userId).Error; err != nil || user.ID == 0 {
		return user, err
	} else {
		return user, nil
	}

}

func (repository *UserRepositoryDB) Update(userId uint, updateUser entities.User) (entities.User, error) {
	user := entities.User{}

	if err := repository.db.First(&user, "id=?", userId).Error; err != nil || user.ID == 0 {
		return user, err
	} else {
		repository.db.Model(&user).Updates(updateUser)
		return user, nil
	}

}
