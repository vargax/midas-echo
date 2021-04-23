package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/controllers"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/repository"
	"gitlab.activarsas.net/cvargasc/midas-echo/env"
	"net/http"
	"os"
	"strconv"
)

func Init() {

	repository.InitRepository()

	e := InitFramework()
	controllers.InitRoutes(e)

	e.Logger.Fatal(e.Start(os.Getenv(env.EchoPort)))
}

func InitFramework() *echo.Echo {
	e := echo.New()
	e.Debug, _ = strconv.ParseBool(os.Getenv(env.DebugEcho))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv(env.CorsOrigin)},
		AllowHeaders: []string{"Authorization"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.Validator = &DataValidator{
		validator: validator.New(),
	}

	return e
}

type DataValidator struct {
	validator *validator.Validate
}

func (dv *DataValidator) Validate(i interface{}) error {
	return dv.validator.Struct(i)
}
