package echo

import "github.com/go-playground/validator/v10"

// DataValidator https://echo.labstack.com/guide/request/#Validate
type DataValidator struct {
	validator *validator.Validate
}

func (dv *DataValidator) Validate(i interface{}) error {
	return dv.validator.Struct(i)
}

func newValidator() *DataValidator {
	return &DataValidator{
		validator: validator.New(),
	}
}

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
