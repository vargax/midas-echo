package echo

import (
	ecb "github.com/labstack/echo-contrib/casbin"
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	"github.com/vargax/midas-echo/env"
	"github.com/vargax/midas-echo/src/echo/auth"
	"github.com/vargax/midas-echo/src/echo/validator"
	"net/http"
	"os"
	"strconv"
)

var e *echo.Echo

func Init() {
	e = echo.New()

	debug, _ := strconv.ParseBool(os.Getenv(env.DebugEcho))
	e.Debug = debug

	e.Use(emw.Logger())
	e.Use(emw.Recover())
	e.Use(emw.Gzip())
	e.Use(emw.Secure())

	cors := os.Getenv(env.CorsOrigin)
	e.Use(emw.CORSWithConfig(emw.CORSConfig{
		AllowOrigins: []string{cors},
		AllowHeaders: []string{"Authorization"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.Use(emw.JWTWithConfig(auth.AuthenticationConfig()))
	e.Use(ecb.MiddlewareWithConfig(auth.AuthorizationConfig()))

	e.Validator = validator.NewValidator()

	Routes(e)

	port := os.Getenv(env.EchoPort)
	e.Logger.Fatal(e.Start(port))
}
