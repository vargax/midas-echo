package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

const (
	echoPort = "ECHO_PORT"

	debugRepo = "DEBUG_REPOSITORY"
	debugJwt  = "DEBUG_JWT"

	jwtIss      = "JWT_ISS"
	jwtAudience = "JWT_AUDIENCE"

	corsOrigin = "CORS_ORIGIN"

	dbHost = "POSTGRES_HOST"
	dbPort = "POSTGRES_PORT"
	dbUser = "POSTGRES_USER"
	dbPass = "POSTGRES_PASSWORD"
	dbName = "POSTGRES_DB"
)

var e *echo.Echo

func Init() {

	EchoInit()
	RepositoryInit()
	AuthMiddlewareInit()
	ControllerInit()

	e.Logger.Fatal(e.Start(os.Getenv(echoPort)))
}

func EchoInit() {
	e = echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv(corsOrigin)},
		AllowHeaders: []string{"Authorization"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.Validator = &DataValidator{
		validator: validator.New(),
	}

}

type DataValidator struct {
	validator *validator.Validate
}

func (dv *DataValidator) Validate(i interface{}) error {
	return dv.validator.Struct(i)
}
