package validator

import (
	"github.com/go-playground/validator/v10"
)

var instance = validator.New()

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value any    `json:"value"`
}

func Struct(data any) []ValidationError {
	errors := []ValidationError{}

	if err := instance.Struct(data); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			verr := ValidationError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Param(),
			}
			errors = append(errors, verr)
		}
	}

	return errors
}
