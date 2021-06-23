package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/vargax/midas-echo/src/middleware"
	"github.com/vargax/midas-echo/src/models"
	"github.com/vargax/midas-echo/src/services"
	"net/http"
)

func PostAppUsers(c echo.Context) error {

	userPost := new(models.PostAppUsers)
	if err := c.Bind(userPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(userPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := services.NewUser(userPost)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, response)
}

func PostPublicTokens(c echo.Context) error {

	tokenPost := new(models.PostPublicToken)
	if err := c.Bind(tokenPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(tokenPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := middleware.JwtTokenFactory(tokenPost)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
