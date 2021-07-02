package main

import (
	"github.com/joho/godotenv"
	"github.com/vargax/midas-echo/src/echo"
	"github.com/vargax/midas-echo/src/postgres"
)

func main() {
	err := godotenv.Load("env/.env")
	if err != nil {
		panic(err)
	}

	postgres.Init()
	echo.Init()
}
