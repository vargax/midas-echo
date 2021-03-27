package main

func HandleNuevoCatalogo(cp *CatalogoPost) (Catalogo, error) {
	nuevoCatalogo := Catalogo{
		EsPublico: cp.EsPublico,
	}
	err := CreateCatalogo(&nuevoCatalogo)

	return nuevoCatalogo, err
}
