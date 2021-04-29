package controllers

import (
	"github.com/labstack/echo/v4"
)

const (
	// App
	app    = "/app"
	tokens = "/tokens"
	users  = "/users"

	// Catalogos
	catalogos = "/catalogos"
	lotes     = "/lotes"

	// Params
	preload    = "preload"
	catalogoId = "catalogoId"
)

var e *echo.Echo

func Routes(framework *echo.Echo) {

	e = framework

	// App
	ag := e.Group(app)
	ag.POST(tokens, PostAppTokens)
	ag.POST(users, PostAppUsers)

	// Catalogos
	cg := e.Group(catalogos)
	cg.GET("", GetCatalogos)
	cg.GET("/:"+catalogoId, GetCatalogosId)

	cg.POST("", PostCatalogos)
	cg.POST("/:"+catalogoId+lotes, PostCatalogosLotes)
}
