package echo

import (
	ecb "github.com/labstack/echo-contrib/casbin"
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	"github.com/vargax/midas-echo"
	"github.com/vargax/midas-echo/env"
	"net/http"
	"os"
	"strconv"
)

var e *echo.Echo
var ss *midas.Services

func Start(s *midas.Services) {
	ss = s

	e = echo.New()

	debug, _ := strconv.ParseBool(os.Getenv(env.DebugEcho))
	e.Debug = debug

	e.Use(emw.Logger())
	//e.Use(emw.Recover())
	e.Use(emw.Gzip())
	e.Use(emw.Secure())

	cors := os.Getenv(env.CorsOrigin)
	e.Use(emw.CORSWithConfig(emw.CORSConfig{
		AllowOrigins: []string{cors},
		AllowHeaders: []string{"Authorization"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.Use(emw.JWTWithConfig(authenticationConfig()))
	e.Use(ecb.MiddlewareWithConfig(authorizationConfig()))

	e.Validator = newValidator()

	routes(e)

	port := os.Getenv(env.EchoPort)
	e.Logger.Fatal(e.Start(port))
}
