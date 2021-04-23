package models

// Domain ******************
// To validate business logic
// *************************

type PostCatalogos struct {
	EsPublico bool `validate:"required"`
}

type PostCatalogosLotes struct {
	Descripcion string `validate:"required"`
}

// App *********************
// To validate Application logic
// *************************

type PostAppToken struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
