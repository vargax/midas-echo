package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/vargax/midas-echo/src/echo/auth"
	"github.com/vargax/midas-echo/src/echo/validator"
	"github.com/vargax/midas-echo/src/postgres"
	"net/http"
)

func PostAppUsers(c echo.Context) error {

	userPost := new(validator.PostAppUsers)
	if err := c.Bind(userPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(userPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := postgres.NewUser(userPost)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, response)
}

func PostAppTokens(c echo.Context) error {

	tokenPost := new(validator.PostPublicToken)
	if err := c.Bind(tokenPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(tokenPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := auth.JwtTokenFactory(tokenPost)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
