package form

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var Validate *validator.Validate

func InitValidate() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func MapValidationError(err validator.FieldError) (string, error) {
	switch err.Tag() {
	case "required":
		return "This field is required", nil
	case "email":
		return "Input valid email address", nil
	case "lte":
		return fmt.Sprintf("Input value should less than or equal %v", err.Param()), nil
	case "gte":
		return fmt.Sprintf("Input value should greater than or equal %v", err.Param()), nil
	}
	return "", errors.New("map error not found")
}
