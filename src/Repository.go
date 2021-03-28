package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
)

var db *gorm.DB

func RepositoryInit() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv(dbHost),
		os.Getenv(dbPort),
		os.Getenv(dbUser),
		os.Getenv(dbName),
		os.Getenv(dbPass),
	)

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	debug, _ := strconv.ParseBool(os.Getenv(debugRepo))
	if debug {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	err = db.AutoMigrate(&Catalogo{}, &Lote{}, &Publicacion{}, &Archivo{})
	if err != nil {
		panic(err)
	}
}

func CreateCatalogo(catalogo *Catalogo) error {
	result := db.Create(catalogo)
	return result.Error
}

func ReadCatalogo(idCatalogo int) (Catalogo, error) {
	var catalogo Catalogo
	result := db.First(&catalogo, idCatalogo)
	return catalogo, result.Error
}

func ReadCatalogos() ([]Catalogo, error) {
	var catalogos []Catalogo
	result := db.Find(&catalogos)

	return catalogos, result.Error
}

func CreateLote(lote *Lote) error {
	result := db.Create(lote)
	return result.Error
}
