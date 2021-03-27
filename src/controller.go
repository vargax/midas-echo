package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ControllerInit() {
	e.GET("/catalogos", GetCatalogos)
}

func GetCatalogos(c echo.Context) error {
	catalogos, err := ReadCatalogos()
	if err != nil {
		e.Logger.Error(err)
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, catalogos)
}
