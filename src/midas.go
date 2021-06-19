package src

import (
	"github.com/go-playground/validator/v10"
	ecb "github.com/labstack/echo-contrib/casbin"
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	"gitlab.activarsas.net/cvargasc/midas-echo/env"
	"gitlab.activarsas.net/cvargasc/midas-echo/src/controllers"
	"gitlab.activarsas.net/cvargasc/midas-echo/src/middleware"
	"gitlab.activarsas.net/cvargasc/midas-echo/src/repository"
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

	e := initEcho()
	controllers.Routes(e)

	e.Logger.Fatal(e.Start(port))
}

func initEcho() *echo.Echo {
	e := echo.New()
	e.Debug = debug

	e.Use(emw.Logger())
	e.Use(emw.Recover())
	e.Use(emw.Gzip())
	e.Use(emw.Secure())

	e.Use(emw.CORSWithConfig(emw.CORSConfig{
		AllowOrigins: []string{cors},
		AllowHeaders: []string{"Authorization"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.Use(emw.JWTWithConfig(middleware.AuthenticationConfig()))
	e.Use(ecb.MiddlewareWithConfig(middleware.AuthorizationConfig()))

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
