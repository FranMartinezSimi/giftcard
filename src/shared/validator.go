package shared

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func ValidateStruct(s interface{}) []*ValidationError {
	var errors []*ValidationError
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationError
			element.Field = err.Field()
			element.Error = formatErrorMessage(err)
			errors = append(errors, &element)
		}
	}
	return errors
}

func formatErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return fmt.Sprintf("This field must be at least %s characters long", err.Param())
	case "max":
		return fmt.Sprintf("This field must be at most %s characters long", err.Param())
	case "gtfield":
		return fmt.Sprintf("This field must be greater than the %s field", err.Param())
	case "gt":
		return fmt.Sprintf("This field must be greater than %s", err.Param())
	case "lt":
		return fmt.Sprintf("This field must be less than %s", err.Param())
		// Add more custom messages for other tags as needed
	default:
		return fmt.Sprintf("Validation failed on field '%s' with tag '%s'", err.Field(), err.Tag())
	}
}

// Helper to return a more structured error for API responses
func FormatValidationErrors(errors []*ValidationError) map[string]interface{} {
	errMap := make(map[string]interface{})
	for _, err := range errors {
		// To avoid overwriting errors if a field has multiple validation failures,
		// we can append to a list of errors for that field or ensure unique keys.
		// For simplicity here, we'll just take the first error for a field.
		// A more robust solution might involve err.Namespace() or a list of messages per field.
		if _, ok := errMap[strings.ToLower(err.Field)]; !ok {
			errMap[strings.ToLower(err.Field)] = err.Error
		}
	}
	return map[string]interface{}{"validation_errors": errMap}
}
