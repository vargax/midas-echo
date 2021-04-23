package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/models"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/repository"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/services"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetCatalogosId(c echo.Context) error {
	var err error
	var idCatalogo int

	idCatalogo, err = strconv.Atoi(c.Param(catalogoId))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	catalogo, err := repository.ReadCatalogo(idCatalogo)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, catalogo)
}

func GetCatalogos(c echo.Context) error {
	preload, _ := strconv.ParseBool(c.QueryParam(preload))

	catalogos, err := repository.ReadCatalogos(preload)
	if err != nil {
		e.Logger.Error(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, catalogos)
}

func PostCatalogos(c echo.Context) error {
	var err error

	catalogoPost := new(models.CatalogoPost)
	if err = c.Bind(catalogoPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(catalogoPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	catalogo, err := services.NewCatalogo(catalogoPost)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, catalogo)
}

func PostCatalogosLotes(c echo.Context) error {
	var err error
	var idCatalogo int

	idCatalogo, err = strconv.Atoi(c.Param(catalogoId))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	lotePost := new(models.LotePost)
	if err = c.Bind(lotePost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(lotePost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	lote, err := services.NewLote(idCatalogo, lotePost)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, lote)
}
