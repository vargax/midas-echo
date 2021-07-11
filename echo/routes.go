package echo

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/vargax/midas-echo"
	"net/http"
	"strconv"
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

func routes(framework *echo.Echo) {

	e = framework

	// App
	ag := e.Group(app)
	ag.POST(users, postAppUsers)
	ag.POST(tokens, postAppTokens)

	// All
	cg := e.Group(catalogos)
	cg.GET("", getCatalogos)
	cg.GET("/:"+catalogoId, getCatalogosId)

	cg.POST("", postCatalogos)
	cg.POST("/:"+catalogoId+lotes, postCatalogosLotes)
}

func postAppUsers(c echo.Context) error {
	userPost := new(PostAppUsers)
	if err := c.Bind(userPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(userPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	usr := midas.User{
		Username: userPost.Username,
		Password: userPost.Password,
		Role:     midas.Role(userPost.Role),
	}

	err := ss.UserSrv.New(&usr)
	if errors.Is(err, midas.UserAlreadyRegistered) {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &usr)
}
func postAppTokens(c echo.Context) error {
	tokenPost := new(PostPublicToken)
	if err := c.Bind(tokenPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(tokenPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := jwtTokenFactory(tokenPost)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func getCatalogos(c echo.Context) error {
	p, _ := strconv.ParseBool(c.QueryParam(preload))

	cc, err := ss.CatSrv.All(p)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, cc)
}
func getCatalogosId(c echo.Context) error {

	idCatalogo, err := strconv.ParseUint(c.Param(catalogoId), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	catalogo, err := ss.CatSrv.Catalogo(uint(idCatalogo))
	if errors.Is(err, midas.RecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, catalogo)
}

func postCatalogos(c echo.Context) error {
	catalogoPost := new(PostCatalogos)
	if err := c.Bind(catalogoPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(catalogoPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctg := midas.Catalogo{
		EsPublico: catalogoPost.EsPublico,
	}
	if err := ss.CatSrv.New(&ctg); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &ctg)
}
func postCatalogosLotes(c echo.Context) error {
	idCatalogo, err := strconv.ParseUint(c.Param(catalogoId), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	lotePost := new(PostCatalogosLotes)
	if err = c.Bind(lotePost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(lotePost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	lt := midas.Lote{
		CatalogoID:  uint(idCatalogo),
		Descripcion: lotePost.Descripcion,
	}
	if err := ss.CatSrv.AddLote(&lt); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &lt)
}
