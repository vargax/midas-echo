package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

const (
	echoPort = "ECHO_PORT"

	debugRepo = "DEBUG_REPOSITORY"

	dbHost = "POSTGRES_HOST"
	dbPort = "POSTGRES_PORT"
	dbUser = "POSTGRES_USER"
	dbPass = "POSTGRES_PASSWORD"
	dbName = "POSTGRES_DB"
)

var e *echo.Echo

func AppInit() {
	e = echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())

}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	AppInit()
	RepositoryInit()
	ControllerInit()

	e.Logger.Fatal(e.Start(os.Getenv(echoPort)))
}
