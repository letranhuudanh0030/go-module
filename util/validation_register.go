package util

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"todo/config"

	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

func getMessageCode(err ErrorResponse) string {
	var msgCode string
	param := strings.Split(err.Param, " ")
	switch err.Tag {
	case "required", "required_if":
		msgCode = config.REQUIRED
	case "min", "min_if":
		msgCode = config.MIN_LENGTH + "_" + param[len(param)-1]
	case "max":
		msgCode = config.MAX_LENGTH + "_" + param[len(param)-1]
	case "datetime":
		msgCode = config.FORMAT_DATE
	case "lttimefield", "ltetimefield", "gttimefield", "gtetimefield", "eqtimefield":
		msgCode = config.TIME_ERROR
	case "valid_input":
		msgCode = config.VALID_LANGUAGE + "_" + param[len(param)-1]
	default:
		msgCode = err.Tag
	}

	return msgCode
}

func registerValidations(v *validator.Validate) error {
	// Time Field Less Than Another Time Field
	if err := v.RegisterValidation("lttimefield", lttimefield); err != nil {
		return err
	}
	// Time Field Less Than or Equal To Another Time Field
	if err := v.RegisterValidation("ltetimefield", ltetimefield); err != nil {
		return err
	}
	// Time Field Greater Than Another Time Field
	if err := v.RegisterValidation("gttimefield", gttimefield); err != nil {
		return err
	}
	// Time Field Greater Than or Equal To Another Time Field
	if err := v.RegisterValidation("gtetimefield", gtetimefield); err != nil {
		return err
	}
	// Time Field Equal To Another Time Field
	if err := v.RegisterValidation("eqtimefield", eqtimefield); err != nil {
		return err
	}
	// Check min_if for array int, string
	if err := v.RegisterValidation("min_if", minIf); err != nil {
		return err
	}
	// Check input english, japanese, vietnamese
	if err := v.RegisterValidation("valid_input", validInput); err != nil {
		return err
	}

	return nil
}

// Check input english, japanese, vietnamese
func validInput(fl validator.FieldLevel) bool {
	text := fl.Field().String()
	tagParts := strings.Split(fl.Param(), " ")
	if len(tagParts) != 1 {
		return false
	}

	lang := tagParts[0]

	// Regular expression patterns for allowed characters
	englishPattern := "^[a-zA-Z0-9 ]*$"
	japanesePattern := "^[ぁ-んァ-ヾ一-龯]*$"
	vietnamesePattern := "^[a-zA-Z0-9À-Ỹà-ỹĂăÂâĐđÊêÔôƠơƯư ]*$"

	var pattern string
	switch lang {
	case "vi":
		pattern = vietnamesePattern
	case "en":
		pattern = englishPattern
	case "ja":
		pattern = japanesePattern
	}

	// Check if the input contains only allowed characters
	return (regexp.MustCompile(pattern).MatchString(text))
}

// Check min_if for array int, string
func minIf(fl validator.FieldLevel) bool {
	field := fl.Field()

	tagParts := strings.Split(fl.Param(), " ")
	if len(tagParts) != 3 {
		return false
	}

	conditionFieldName := tagParts[0]
	expectedCondition := tagParts[1]
	expectedMinStr := tagParts[2]

	boolValue, err := strconv.ParseBool(expectedCondition)
	if err != nil {
		return false
	}

	conditionField := fl.Parent().FieldByName(conditionFieldName)

	if conditionField.Kind() == reflect.Bool {
		if boolValue == conditionField.Bool() {
			expectedMin, err := strconv.Atoi(expectedMinStr)
			if err != nil {
				return false
			}
			switch field.Interface().(type) {
			case []string:
				return len(field.Interface().([]string)) >= expectedMin
			case pq.StringArray:
				return len(field.Interface().(pq.StringArray)) >= expectedMin
			case []int:
				return len(field.Interface().([]int)) >= expectedMin
			case pq.Int64Array:
				return len(field.Interface().(pq.Int64Array)) >= expectedMin
			default:
				// Handle other types if needed
				return true
			}
		}
	}

	return true
}

// Time Field Less Than Another Time Field
func lttimefield(fl validator.FieldLevel) bool {
	var fieldValue time.Time
	var compareValue time.Time
	checkTimeRelation(fl, &fieldValue, &compareValue)
	return fieldValue.Before(compareValue)
}

// Time Field Less Than or Equal To Another Time Field
func ltetimefield(fl validator.FieldLevel) bool {
	var fieldValue time.Time
	var compareValue time.Time
	checkTimeRelation(fl, &fieldValue, &compareValue)
	return fieldValue.Before(compareValue) || fieldValue.Equal(compareValue)
}

// Time Field Greater Than Another Time Field
func gttimefield(fl validator.FieldLevel) bool {
	var fieldValue time.Time
	var compareValue time.Time
	checkTimeRelation(fl, &fieldValue, &compareValue)
	return fieldValue.After(compareValue)
}

// Time Field Greater Than or Equal To Another Time Field
func gtetimefield(fl validator.FieldLevel) bool {
	var fieldValue time.Time
	var compareValue time.Time
	checkTimeRelation(fl, &fieldValue, &compareValue)
	return fieldValue.After(compareValue) || fieldValue.Equal(compareValue)
}

// Time Field Equal To Another Time Field
func eqtimefield(fl validator.FieldLevel) bool {
	var fieldValue time.Time
	var compareValue time.Time
	checkTimeRelation(fl, &fieldValue, &compareValue)
	return fieldValue.Equal(compareValue)
}

func checkTimeRelation(fl validator.FieldLevel, fieldValue *time.Time, compareValue *time.Time) {
	fieldName := fl.Param() // Get the field name to compare with
	*fieldValue = fl.Field().Interface().(time.Time)
	structValue := fl.Parent()

	// Retrieve the value of the field to compare with
	compareField := structValue.FieldByName(fieldName)
	*compareValue = compareField.Interface().(time.Time)
}

func ConvertFieldName(elem *ErrorResponse) {
	namespace := elem.Namespace
	fileName := elem.Field
	// Define a regular expression pattern to match indices within brackets
	pattern := `\[([0-9]+)\]`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find all matches of the pattern in the input string
	matches := re.FindAllStringSubmatch(namespace, -1)
	if len(matches) > 0 {
		// Extract and convert indices from the matches
		var index string
		for _, match := range matches {
			indexStr := match[1] // Submatch containing the index
			index += "_" + indexStr
		}

		// Extract the field name (last part)
		parts := strings.Split(namespace, ".")
		namespace := parts[len(parts)-1]

		// Construct the desired fileName string
		fileName = fmt.Sprintf("%s%s", namespace, index)
	}

	elem.Field = fileName
}
