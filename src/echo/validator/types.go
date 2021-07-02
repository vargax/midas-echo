package validator

// Domain ******************
// To validate business logic
// *************************

type (
	PostCatalogos struct {
		EsPublico bool
	}

	PostCatalogosLotes struct {
		Descripcion string `validate:"required"`
	}
)

// App *********************
// To validate Application logic
// *************************

type (
	PostAppUsers struct {
		Username string `json:"username" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required,oneof=admin user"`
	}

	PostPublicToken struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)
