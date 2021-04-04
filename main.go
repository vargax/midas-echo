package main

import (
	"github.com/joho/godotenv"
	"gitlab.activarsas.net/cvargasc/midas-echo/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app.Init()

}
