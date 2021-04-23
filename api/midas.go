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

var (
	debug bool
	port  string
	cors  string
)

func Env() {
	debug, _ = strconv.ParseBool(os.Getenv(env.DebugEcho))
	port = os.Getenv(env.EchoPort)
	cors = os.Getenv(env.CorsOrigin)
}

func Init() {
	repository.Init()

	e := InitEcho()
	controllers.Routes(e)

	e.Logger.Fatal(e.Start(port))
}

func InitEcho() *echo.Echo {
	e := echo.New()
	e.Debug = debug

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{cors},
		AllowHeaders: []string{"Authorization"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.Validator = &DataValidator{
		validator: validator.New(),
	}

	return e
}

// Validator ***************
// https://echo.labstack.com/guide/request/#Validate
// *************************

type DataValidator struct {
	validator *validator.Validate
}

func (dv *DataValidator) Validate(i interface{}) error {
	return dv.validator.Struct(i)
}
