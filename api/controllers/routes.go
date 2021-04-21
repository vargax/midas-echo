package controllers

import (
	"github.com/labstack/echo/v4"
)

const (
	catalogosPath = "/catalogos"
	lotesPath     = "/lotes"

	PreloadParam    = "preload"
	CatalogoIdParam = "catalogoId"
)

var e *echo.Echo

func InitRoutes(framework *echo.Echo) {

	e = framework

	// Catalogos
	e.POST(catalogosPath, PostCatalogos)

	e.GET(catalogosPath, GetCatalogos)
	e.GET(catalogosPath+"/:"+CatalogoIdParam, GetCatalogosId)

	e.POST(catalogosPath+"/:"+CatalogoIdParam+lotesPath, PostCatalogosLotes)
}
