package services

import (
	"gitlab.activarsas.net/cvargasc/midas-echo/api/models"
	"gitlab.activarsas.net/cvargasc/midas-echo/api/repository"
)

func NewCatalogo(cp *models.CatalogoPost) (models.Catalogo, error) {
	nuevoCatalogo := models.Catalogo{
		EsPublico: cp.EsPublico,
	}
	err := repository.CreateCatalogo(&nuevoCatalogo)

	return nuevoCatalogo, err
}

func NewLote(idCatalogo int, lp *models.LotePost) (models.Lote, error) {
	nuevoLote := models.Lote{
		CatalogoID:  idCatalogo,
		Descripcion: lp.Descripcion,
	}
	err := repository.CreateLote(&nuevoLote)

	return nuevoLote, err
}
