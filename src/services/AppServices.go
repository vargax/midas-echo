package services

import (
	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"github.com/vargax/midas-echo/src/models"
	"github.com/vargax/midas-echo/src/repository"
	"github.com/vargax/midas-echo/src/utils"
	"net/http"
)

func NewUser(userToBe *models.PostAppUsers) (*models.User, error) {
	pass, err := utils.EncodePassword(userToBe.Password)
	if err != nil {
		return nil, err
	}

	newUser := models.User{
		Username: userToBe.Username,
		Password: pass,
		Role:     userToBe.Role,
	}

	err = repository.CreateUser(&newUser)
	if e, ok := err.(*pgconn.PgError); ok && e.Code == repository.DuplicateKey {
		return nil, echo.NewHTTPError(http.StatusConflict, e.Detail)
	}

	return &newUser, err
}
