package main

import (
	"github.com/joho/godotenv"
	"gitlab.activarsas.net/cvargasc/midas-echo/api"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/repository"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/services"
)

func main() {
	err := godotenv.Load("env/.env")
	if err != nil {
		panic(err)
	}

	repository.Env()
	services.Env()
	api.Env()

	api.Init()
}
