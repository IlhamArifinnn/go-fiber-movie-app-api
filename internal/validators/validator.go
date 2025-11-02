package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Init() {
	validate = validator.New()
}

var validationErrorsMessages = map[string]string{
	"required": "This field is required",
	"email":    "Invalid email format",
	"min":      "Value is below the minimum allowed",
	"max":      "Value exceeds the maximum allowed",
	"dive":     "Invalid value in the list",
	"url":      "Invalid URL format",
	"datetime": "Invalid date format, expected YYYY-MM-DD",
	"numeric":  "This field must be a numeric value",
}

func ValidateStruct(s interface{}) []string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()
		tag := err.Tag()

		msg, exists := validationErrorsMessages[tag]
		if !exists {
			msg = fmt.Sprintf("Validation failed on %s rule", tag)
		}
		errors = append(errors, fmt.Sprintf("%s: %s", fieldName, msg))
	}
	return errors
}
