package repository

import (
	"github.com/vargax/midas-echo/src/models"
)

func CreateUser(user *models.User) error {
	result := db.Create(user)
	return result.Error
}

func ReadUser(username string) (models.User, error) {
	user := models.User{
		Username: username,
	}
	result := db.Where(&user).First(&user)

	return user, result.Error
}
