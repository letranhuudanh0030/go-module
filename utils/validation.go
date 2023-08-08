package utils

import (
	"reflect"
	"strings"
	"todo/config"

	"github.com/go-playground/validator/v10"
)

type (
	XValidator struct {
		validator *validator.Validate
	}

	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
		KeyJson     string
	}
)

var validate = validator.New()

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})

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

func getMessageCode(tag string) string {
	var code string

	switch tag {
	case "required":
		code = config.REQUIRED
	case "min":
		code = config.MIN_LENGTH
	case "max":
		code = config.MAX_LENGTH
	}

	return code
}

func Validator(data interface{}) map[string]string {
	myValidator := &XValidator{
		validator: validate,
	}

	// Validation
	if errs := myValidator.Validate(data); len(errs) > 0 && errs[0].Error {
		errMsgs := make(map[string]string)

		for _, err := range errs {
			errMsgs[err.FailedField] = getMessageCode(err.Tag)
		}

		return errMsgs
	}

	return nil
}
