package controllers

import (
	"github.com/labstack/echo/v4"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/models"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/services"
	"net/http"
)

func PostAppToken(c echo.Context) error {
	var err error

	tokenPost := new(models.PostAppToken)
	if err = c.Bind(tokenPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(tokenPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := services.NewJwtToken(tokenPost)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
