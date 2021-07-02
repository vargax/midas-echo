package postgres

import (
	"fmt"
	"github.com/vargax/midas-echo/env"
	"github.com/vargax/midas-echo/src"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
)

const (
	DuplicateKey = "23505"
)

var db *gorm.DB

func Init() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv(env.DbHost),
		os.Getenv(env.DbPort),
		os.Getenv(env.DbUser),
		os.Getenv(env.DbName),
		os.Getenv(env.DbPass),
	)

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	debug, _ := strconv.ParseBool(os.Getenv(env.DebugRepo))
	if debug {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	err = db.AutoMigrate(
		&src.Catalogo{}, &src.Lote{}, &src.Publicacion{},
		&src.File{}, &src.User{})
	if err != nil {
		panic(err)
	}

	err = initDB()
	if err != nil {
		panic(err)
	}
}

func initDB() error {
	count := int64(-1)

	result := db.Model(&src.User{}).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	if count == 0 {
		defaultUser := os.Getenv(env.DefaultUser)
		defaultPass := os.Getenv(env.DefaultPass)

		fmt.Printf("Creating default user %s with password %s", defaultUser, defaultPass)

		defaultPass, err := EncodePassword(defaultPass)
		if err != nil {
			return err
		}

		admin := src.User{
			Username: defaultUser,
			Password: defaultPass,
			Role:     src.RoleAdmin,
		}

		result := db.Create(&admin)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
