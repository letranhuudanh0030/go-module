package util

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"text/template"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CopyStruct(source interface{}, target interface{}) error {
	sourceValue := reflect.ValueOf(source)
	targetType := reflect.TypeOf(target)

	if targetType.Kind() != reflect.Ptr || targetType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to a struct")
	}

	targetValue := reflect.ValueOf(target).Elem()

	for i := 0; i < targetValue.NumField(); i++ {
		fieldName := targetValue.Type().Field(i).Name
		sourceField := sourceValue.Elem().FieldByName(fieldName)
		if sourceField.IsValid() {
			targetValue.Field(i).Set(sourceField)
		}
	}

	return nil
}

func GetTransactionFromCtx(c *fiber.Ctx) (*gorm.DB, bool) {
	// Access the transaction from the Fiber context
	ctx := c.Context().UserValue("TransactionContextKey")
	if ctx != nil {
		tx, ok := ctx.(context.Context).Value("TransactionContextKey").(*gorm.DB)
		if ok {
			return tx, true
		}
	}
	return nil, false
}

func ParseTemplateString(name string, str string, data interface{}) (string, error) {
	strTmpl, err := template.New(name).Parse(str)
	if err != nil {
		return "", err
	}

	var templateOutput bytes.Buffer

	if err := strTmpl.Execute(&templateOutput, data); err != nil {
		return "", err
	}

	return templateOutput.String(), nil
}
