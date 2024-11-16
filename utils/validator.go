package utils

import (
	"fmt"

	"github.com/go-playground/validator"
)

var validate = validator.New()

func ValidateStruct(data interface{}) (map[string]string, bool) {
	if err := validate.Struct(data); err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[err.Field()] = fmt.Sprintf("failed on %s", err.Tag())
		}

		return validationErrors, false
	}
	return nil, true
}
