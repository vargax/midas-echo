package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/vargax/midas-echo/src/middleware"
)

// Public endpoints ********
// All routes under /public skip Authentication and Authorization (services.skipper function)
// *************************
const (
	tokens = "/tokens"
)

// Protected endpoints *****
// All other routes are Protected by default (Authentication and Authorization enforced)
// *************************
const (
	// App
	app   = "/app"
	users = "/users"

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

	// Public endpoints ********
	pg := e.Group(middleware.Public)
	pg.POST(tokens, PostPublicTokens)

	// Protected endpoints *****
	// App
	ag := e.Group(app)
	ag.POST(users, PostAppUsers)

	// Catalogos
	cg := e.Group(catalogos)
	cg.GET("", GetCatalogos)
	cg.GET("/:"+catalogoId, GetCatalogosId)

	cg.POST("", PostCatalogos)
	cg.POST("/:"+catalogoId+lotes, PostCatalogosLotes)
}
