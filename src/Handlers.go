package main

func HandleNuevoCatalogo(cp *CatalogoPost) (Catalogo, error) {
	nuevoCatalogo := Catalogo{
		EsPublico: cp.EsPublico,
	}
	err := CreateCatalogo(&nuevoCatalogo)

	return nuevoCatalogo, err
}

func HandleNuevoLote(idCatalogo int, lp *LotePost) (Lote, error) {
	nuevoLote := Lote{
		CatalogoID:  idCatalogo,
		Descripcion: lp.Descripcion,
	}
	err := CreateLote(&nuevoLote)

	return nuevoLote, err
}
