package main

import (
	"github.com/joho/godotenv"
	"github.com/vargax/midas-echo/src"
	"github.com/vargax/midas-echo/src/middleware"
	"github.com/vargax/midas-echo/src/repository"
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
