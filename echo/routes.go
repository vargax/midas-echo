package echo

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/vargax/midas-echo"
	"github.com/vargax/midas-echo/echo/validator"
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

func Routes(framework *echo.Echo) {

	e = framework

	// App
	ag := e.Group(app)
	ag.POST(users, PostAppUsers)
	ag.POST(tokens, PostAppTokens)

	// All
	cg := e.Group(catalogos)
	cg.GET("", GetCatalogos)
	cg.GET("/:"+catalogoId, GetCatalogosId)

	cg.POST("", PostCatalogos)
	cg.POST("/:"+catalogoId+lotes, PostCatalogosLotes)
}

func PostAppUsers(c echo.Context) error {
	userPost := new(validator.PostAppUsers)
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

	if err := ss.UserSrv.New(&usr); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &usr)
}
func PostAppTokens(c echo.Context) error {
	tokenPost := new(validator.PostPublicToken)
	if err := c.Bind(tokenPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(tokenPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := JwtTokenFactory(tokenPost)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func GetCatalogos(c echo.Context) error {
	p, _ := strconv.ParseBool(c.QueryParam(preload))

	cc, err := ss.CatSrv.All(p)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, cc)
}
func GetCatalogosId(c echo.Context) error {

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

func PostCatalogos(c echo.Context) error {
	catalogoPost := new(validator.PostCatalogos)
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

	lt := midas.Lote{
		CatalogoID:  uint(idCatalogo),
		Descripcion: lotePost.Descripcion,
	}
	if err := ss.CatSrv.AddLote(&lt); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &lt)
}
