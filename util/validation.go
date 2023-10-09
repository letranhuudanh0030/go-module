package util

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type (
	XValidator struct {
		validator *validator.Validate
	}

	ErrorResponse struct {
		Namespace       string `json:"namespace"`
		Field           string `json:"field"`
		StructNamespace string `json:"structNamespace"`
		StructField     string `json:"structField"`
		Tag             string `json:"tag"`
		ActualTag       string `json:"actualTag"`
		Kind            string `json:"kind"`
		Type            string `json:"type"`
		Value           string `json:"value"`
		Param           string `json:"param"`
		Message         string `json:"message"`
		Error           bool
	}
)

var validate = validator.New()

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	// validationErrors := []ErrorResponse{}
	validationErrors := []ErrorResponse{}
	if err := registerValidations(validate); err != nil {
		panic(err)
	}

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
			elem := ErrorResponse{
				Namespace:       err.Namespace(),
				Field:           err.Field(),
				StructNamespace: err.StructNamespace(),
				StructField:     err.StructField(),
				Tag:             err.Tag(),
				ActualTag:       err.ActualTag(),
				Kind:            fmt.Sprintf("%v", err.Kind()),
				Type:            fmt.Sprintf("%v", err.Type()),
				Value:           fmt.Sprintf("%v", err.Value()),
				Param:           err.Param(),
				Message:         err.Error(),
				Error:           true,
			}
			ConvertFieldName(&elem)
			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func Validator(data interface{}) map[string]string {
	myValidator := &XValidator{
		validator: validate,
	}

	// Validation
	if errs := myValidator.Validate(data); len(errs) > 0 && errs[0].Error {
		errMsgs := make(map[string]string)

		for _, err := range errs {
			errMsgs[err.Field] = getMessageCode(err)
		}

		return errMsgs
	}

	return nil
}
