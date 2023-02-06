package validation

import "github.com/go-playground/validator"

type ValidatinResponse struct {
	Field      string
	Validation string
}

func ValidateMessage(err error) []ValidatinResponse {
	var validationMessage []ValidatinResponse
	for _, err := range err.(validator.ValidationErrors) {
		validationMessage = append(validationMessage, ValidatinResponse{
			Field:      err.Field(),
			Validation: err.ActualTag(),
		})
	}
	return validationMessage
}
