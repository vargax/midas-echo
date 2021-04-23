package repository

import (
	"fmt"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/models"
	"gitlab.activarsas.net/cvargasc/midas-echo/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
)

var dsn string

func Env() {
	dsn = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv(env.DbHost),
		os.Getenv(env.DbPort),
		os.Getenv(env.DbUser),
		os.Getenv(env.DbName),
		os.Getenv(env.DbPass),
	)
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

	err = db.AutoMigrate(&models.Catalogo{}, &models.Lote{}, &models.Publicacion{}, &models.File{})
	if err != nil {
		panic(err)
	}
}

func CreateCatalogo(catalogo *models.Catalogo) error {
	result := db.Create(catalogo)
	return result.Error
}

func ReadCatalogo(idCatalogo int) (models.Catalogo, error) {
	var catalogo models.Catalogo
	result := db.Preload(clause.Associations).First(&catalogo, idCatalogo)
	return catalogo, result.Error
}

func ReadCatalogos(preload bool) ([]models.Catalogo, error) {
	var catalogos []models.Catalogo
	var result *gorm.DB

	if preload {
		result = db.Preload(clause.Associations).Find(&catalogos)
	} else {
		result = db.Find(&catalogos)
	}

	return catalogos, result.Error
}

func CreateLote(lote *models.Lote) error {
	result := db.Create(lote)
	return result.Error
}
