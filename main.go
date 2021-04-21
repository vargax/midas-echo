package main

import (
	"github.com/joho/godotenv"
	"gitlab.activarsas.net/cvargasc/midas-echo/api"
)

func main() {
	err := godotenv.Load("env/.env")
	if err != nil {
		panic(err)
	}

	api.Init()

}
