package app

type CatalogoPost struct {
	EsPublico bool `validate:"required"`
}

type LotePost struct {
	Descripcion string `validate:"required"`
}
