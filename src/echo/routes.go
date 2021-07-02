package echo

import (
	"github.com/labstack/echo/v4"
)

const (
	// App
	app    = "/app"
	users  = "/users"
	tokens = "/tokens"

	// Catalogos
	catalogos = "/catalogos"
	lotes     = "/lotes"

	// Params
	preload    = "preload"
	catalogoId = "catalogoId"
)

func Routes(framework *echo.Echo) {

	e = framework

	// App
	ag := e.Group(app)
	ag.POST(users, PostAppUsers)
	ag.POST(tokens, PostAppTokens)

	// Catalogos
	cg := e.Group(catalogos)
	cg.GET("", GetCatalogos)
	cg.GET("/:"+catalogoId, GetCatalogosId)

	cg.POST("", PostCatalogos)
	cg.POST("/:"+catalogoId+lotes, PostCatalogosLotes)
}
