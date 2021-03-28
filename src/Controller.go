package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

const (
	catalogosPath = "/catalogos"
	lotesPath     = "/lotes"

	catalogoIdParam = "catalogoId"
)

func ControllerInit() {
	e.POST(catalogosPath, PostCatalogos)
	e.GET(catalogosPath, GetCatalogos)
	e.GET(catalogosPath+"/:"+catalogoIdParam, GetCatalogosId)

	e.POST(catalogosPath+"/:"+catalogoIdParam+lotesPath, PostCatalogosLotes)
}

func GetCatalogosId(c echo.Context) error {
	var err error
	var idCatalogo int

	idCatalogo, err = strconv.Atoi(c.Param(catalogoIdParam))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	catalogo, err := ReadCatalogo(idCatalogo)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, catalogo)
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
	var err error

	catalogoPost := new(CatalogoPost)
	if err = c.Bind(catalogoPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(catalogoPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	catalogo, err := HandleNuevoCatalogo(catalogoPost)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, catalogo)
}

func PostCatalogosLotes(c echo.Context) error {
	var err error
	var idCatalogo int

	idCatalogo, err = strconv.Atoi(c.Param(catalogoIdParam))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	lotePost := new(LotePost)
	if err = c.Bind(lotePost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(lotePost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	lote, err := HandleNuevoLote(idCatalogo, lotePost)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, lote)
}
