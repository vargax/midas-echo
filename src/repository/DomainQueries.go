package repository

import (
	"github.com/vargax/midas-echo/src/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateCatalogo(catalogo *models.Catalogo) error {
	result := db.Create(catalogo)
	return result.Error
}

func ReadCatalogo(idCatalogo uint) (models.Catalogo, error) {
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
