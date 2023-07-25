package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type XValidator struct {
	validator *validator.Validate
}

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

var validate = validator.New()

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	var validationErrors []ErrorResponse

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func (v XValidator) ValidateWithMessage(data interface{}) (formattedMsg string, ok bool) {

	if errs := v.Validate(data); len(errs) > 0 && errs[0].Error {
		errMsg := make([]string, 0)
		for _, err := range errs {
			errMsg = append(errMsg, fmt.Sprintf(
				"%s field is invalid",
				err.FailedField,
			))
		}

		return strings.Join(errMsg, ", "), false
	}

	return "", true
}
