package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gitlab.activarsas.net/cvargasc/midas-echo/env"
	"net/http"
	"os"
	"time"
)

var (
	secret      string
	tokenSecret string
	tokenLive   time.Duration
)

func InitAppController() {

	secret = os.Getenv(env.JwtSecret)
	tokenSecret = os.Getenv(env.JwtTokenSecret)

	var err error
	tokenLive, err = time.ParseDuration(os.Getenv(env.JwtTokenExp))
	if err != nil {
		panic(err)
	}
}

func PostAppToken(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(tokenLive).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
