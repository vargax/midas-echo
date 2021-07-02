package validator

import "github.com/go-playground/validator/v10"

// DataValidator https://echo.labstack.com/guide/request/#Validate
type DataValidator struct {
	validator *validator.Validate
}

func NewValidator() *DataValidator {
	return &DataValidator{
		validator: validator.New(),
	}
}

func (dv *DataValidator) Validate(i interface{}) error {
	return dv.validator.Struct(i)
}
