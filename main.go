package main

import (
	"github.com/joho/godotenv"
	"gitlab.activarsas.net/cvargasc/midas-echo/src"
	"gitlab.activarsas.net/cvargasc/midas-echo/src/middleware"
	"gitlab.activarsas.net/cvargasc/midas-echo/src/repository"
)

func main() {
	err := godotenv.Load("env/.env")
	if err != nil {
		panic(err)
	}

	repository.Env()
	middleware.Env()
	src.Env()

	src.Init()
}
