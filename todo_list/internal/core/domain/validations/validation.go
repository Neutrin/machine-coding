package validations

import (
	"fmt"
	"strings"

	validatorLib "github.com/go-playground/validator/v10"
	"github.com/todo_list/internal/core/domain/enums"
)

var validator *validatorLib.Validate

var customTagValidator map[string]func(f1 validatorLib.FieldLevel) bool

func InitValidator() {
	validator = validatorLib.New()
	customTagValidator = map[string]func(f1 validatorLib.FieldLevel) bool{
		"custom_status": func(f1 validatorLib.FieldLevel) bool {
			value := f1.Field().Interface().(enums.Status)
			return strings.Compare(value.String(), "not present") != 0

		},
		"custom_tags": func(f1 validatorLib.FieldLevel) bool {
			value := f1.Field().Interface().(enums.Tags)
			return strings.Compare(value.String(), "not present") != 0
		},
	}
	for curTag, curValidator := range customTagValidator {
		validator.RegisterValidation(curTag, curValidator)
	}
}

// This will return the first field for whch validation failed
func ValidateStruct(input interface{}) error {
	if err := validator.Struct(input); err != nil {
		for _, curError := range err.(validatorLib.ValidationErrors) {
			return fmt.Errorf("validation failed for field %s ", curError.StructField())
		}
	}
	return nil
}
