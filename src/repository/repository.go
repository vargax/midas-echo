package repository

import (
	"fmt"
	"github.com/vargax/midas-echo/env"
	"github.com/vargax/midas-echo/src/models"
	"github.com/vargax/midas-echo/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
)

const (
	DuplicateKey = "23505"
)

var (
	dsn         string
	defaultUser string
	defaultPass string
)

func Env() {
	dsn = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv(env.DbHost),
		os.Getenv(env.DbPort),
		os.Getenv(env.DbUser),
		os.Getenv(env.DbName),
		os.Getenv(env.DbPass),
	)

	defaultUser = os.Getenv(env.DefaultUser)
	defaultPass = os.Getenv(env.DefaultPass)

}

var db *gorm.DB

func Init() {
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
		&models.Catalogo{}, &models.Lote{}, &models.Publicacion{},
		&models.File{}, &models.User{})
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

	result := db.Model(&models.User{}).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	if count == 0 {

		fmt.Printf("Creating default user %s with password %s", defaultUser, defaultPass)

		defaultPass, err := utils.EncodePassword(defaultPass)
		if err != nil {
			return err
		}

		admin := models.User{
			Username: defaultUser,
			Password: defaultPass,
		}

		result := db.Create(&admin)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
