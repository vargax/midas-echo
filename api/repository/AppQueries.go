package repository

import "gitlab.activarsas.net/cvargasc/midas-echo/api/models"

func ReadUser(username string) (models.User, error) {
	user := models.User{
		Username: username,
	}
	result := db.Where(&user).First(&user)

	return user, result.Error
}
