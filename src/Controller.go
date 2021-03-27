package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	uriCatalogos = "/catalogos"
)

func ControllerInit() {
	e.GET(uriCatalogos, GetCatalogos)
	e.POST(uriCatalogos, PostCatalogos)
}

func GetCatalogos(c echo.Context) error {
	catalogos, err := ReadCatalogos()
	if err != nil {
		e.Logger.Error(err)
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, catalogos)
}

func PostCatalogos(c echo.Context) error {
	catalogoPost := new(CatalogoPost)

	err := c.Bind(catalogoPost)
	if err != nil {
		return err
	}
	catalogo, err := HandleNuevoCatalogo(catalogoPost)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, catalogo)
}
