package controllers

import (
	"github.com/labstack/echo/v4"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/models"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/services"
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

func PostAppTokens(c echo.Context) error {

	tokenPost := new(models.PostAppToken)
	if err := c.Bind(tokenPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(tokenPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := services.NewJwtToken(tokenPost)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
