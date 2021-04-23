package models

// Domain ******************
// To validate business logic
// *************************

type CatalogoPost struct {
	EsPublico bool `validate:"required"`
}

type LotePost struct {
	Descripcion string `validate:"required"`
}

// App *********************
// To validate Application logic
// *************************

type TokenPost struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
