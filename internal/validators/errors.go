package validators

import (
	"github.com/dreezy305/library-core-service/internal/types"
	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []types.FieldError {
	var errors []types.FieldError

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return errors
	}

	for _, e := range validationErrors {
		errors = append(errors, types.FieldError{
			Field:   e.Field(),
			Message: e.Tag(), // e.g. "required", "email", "min"
		})
	}

	return errors
}
