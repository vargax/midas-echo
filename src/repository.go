package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
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
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Catalogo{}, &Lote{}, &Publicacion{}, &Archivo{})
}

func ReadCatalogos() ([]Catalogo, error) {
	var catalogos []Catalogo
	result := db.Find(&catalogos)

	return catalogos, result.Error
}
