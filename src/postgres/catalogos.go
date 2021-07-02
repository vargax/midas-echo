package postgres

import (
	"github.com/vargax/midas-echo/src"
	"github.com/vargax/midas-echo/src/echo/validator"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewCatalogo(cp *validator.PostCatalogos) (*src.Catalogo, error) {
	nuevoCatalogo := src.Catalogo{
		EsPublico: cp.EsPublico,
	}
	result := db.Create(&nuevoCatalogo)

	return &nuevoCatalogo, result.Error
}

func Catalogo(idCatalogo uint) (src.Catalogo, error) {
	var catalogo src.Catalogo
	result := db.Preload(clause.Associations).First(&catalogo, idCatalogo)
	return catalogo, result.Error
}

func Catalogos(preload bool) ([]src.Catalogo, error) {
	var catalogos []src.Catalogo
	var result *gorm.DB

	if preload {
		result = db.Preload(clause.Associations).Find(&catalogos)
	} else {
		result = db.Find(&catalogos)
	}

	return catalogos, result.Error
}

func NewLote(idCatalogo uint, lp *validator.PostCatalogosLotes) (*src.Lote, error) {
	nuevoLote := src.Lote{
		CatalogoID:  idCatalogo,
		Descripcion: lp.Descripcion,
	}
	result := db.Create(&nuevoLote)

	return &nuevoLote, result.Error
}
