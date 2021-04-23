package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/models"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/repository"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/utils"
	"gitlab.activarsas.net/cvargasc/midas-echo/env"
	"gorm.io/gorm"
	"os"
	"time"
)

var (
	secret    string
	tokenLive time.Duration
)

func Env() {
	var err error

	secret = os.Getenv(env.JwtSecret)

	tokenLive, err = time.ParseDuration(os.Getenv(env.JwtTokenExp))
	if err != nil {
		panic(err)
	}
}

func NewJwtToken(tp *models.PostAppToken) (*models.JwtToken, error) {

	user, err := repository.ReadUser(tp.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, echo.ErrUnauthorized
	}
	if err != nil {
		return nil, err
	}

	if !utils.PasswordMatch(user.Password, tp.Password) {
		return nil, echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["UserID"] = user.ID
	claims["exp"] = time.Now().Add(tokenLive).Unix()

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	response := models.JwtToken{
		Token: signedToken,
	}
	return &response, nil
}
