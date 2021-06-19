package services

import (
	"gitlab.activarsas.net/cvargasc/midas-echo/src/models"
	"gitlab.activarsas.net/cvargasc/midas-echo/src/repository"
)

func NewCatalogo(cp *models.PostCatalogos) (models.Catalogo, error) {
	nuevoCatalogo := models.Catalogo{
		EsPublico: cp.EsPublico,
	}
	err := repository.CreateCatalogo(&nuevoCatalogo)

	return nuevoCatalogo, err
}

func NewLote(idCatalogo uint, lp *models.PostCatalogosLotes) (models.Lote, error) {
	nuevoLote := models.Lote{
		CatalogoID:  idCatalogo,
		Descripcion: lp.Descripcion,
	}
	err := repository.CreateLote(&nuevoLote)

	return nuevoLote, err
}
