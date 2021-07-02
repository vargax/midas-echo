package echo

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/vargax/midas-echo/src/echo/validator"
	"github.com/vargax/midas-echo/src/postgres"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetCatalogosId(c echo.Context) error {

	idCatalogo, err := strconv.ParseUint(c.Param(catalogoId), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	catalogo, err := postgres.Catalogo(uint(idCatalogo))
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

	catalogos, err := postgres.Catalogos(preload)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, catalogos)
}

func PostCatalogos(c echo.Context) error {

	catalogoPost := new(validator.PostCatalogos)
	if err := c.Bind(catalogoPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(catalogoPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	catalogo, err := postgres.NewCatalogo(catalogoPost)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, catalogo)
}

func PostCatalogosLotes(c echo.Context) error {

	idCatalogo, err := strconv.ParseUint(c.Param(catalogoId), 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	lotePost := new(validator.PostCatalogosLotes)
	if err = c.Bind(lotePost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(lotePost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	lote, err := postgres.NewLote(uint(idCatalogo), lotePost)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, lote)
}
