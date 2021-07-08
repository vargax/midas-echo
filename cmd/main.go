package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vargax/midas-echo"
	"github.com/vargax/midas-echo/echo"
	"github.com/vargax/midas-echo/env"
	"github.com/vargax/midas-echo/gorm"
	"github.com/vargax/midas-echo/srv"
	"os"
)

var (
	storager midas.StorageSrv

	userer    midas.UserSrv
	cataloger midas.CatalogoSrv
)

func main() {
	if err := godotenv.Load("env/.env"); err != nil {
		panic(err)
	}

	storager = gorm.New()
	userer = srv.NewUserSrv(storager)
	cataloger = srv.NewCatalogoSrv(storager)

	// Init DB with default admin
	if storager.DbInitRequired() {
		admin := midas.User{
			Username: os.Getenv(env.DefaultUser),
			Password: os.Getenv(env.DefaultPass),
			Role:     midas.RoleAdmin,
		}

		if err := userer.New(&admin); err != nil {
			panic(err)
		}

		fmt.Printf("Default %s with username %s and password %s created!", admin.Role, admin.Username, admin.Password)
	}

	echo.Start(&midas.Services{UserSrv: userer, CatSrv: cataloger})
}
