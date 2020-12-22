package hw09_struct_validator //nolint:golint,stylecheck
import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

var ErrorNotStruct = errors.New("not struct")
var ErrorNotValidDefinition = errors.New("not valid definition")
var ErrorUnknownRuleKey = errors.New("unknown rule key")

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var resString string
	for _, validationError := range v {
		resString += fmt.Sprintf("error in field: %v -- %v\n", validationError.Field, validationError.Err)
	}
	return resString
}

func Validate(v interface{}) error {
	rValue := reflect.ValueOf(v)

	rKind := rValue.Kind()
	if rKind != reflect.Struct {
		return ErrorNotStruct
	}

	rType := rValue.Type()

	var validationErrors ValidationErrors

	for i := 0; i < rValue.NumField(); i++ {
		valueField := rValue.Field(i)

		if !valueField.CanInterface() {
			continue
		}
		structFieldParam := rType.Field(i)

		validateTags := structFieldParam.Tag.Get("validate")
		if validateTags == "" {
			continue
		}

		rules := strings.Split(validateTags, "|")
		for _, rule := range rules {
			ruleKey, ruleValue, err := getRule(rule)
			if err != nil {
				return err
			}

			validator := getValidator(ruleKey)
			if validator == nil {
				return ErrorUnknownRuleKey
			}

			err = validator.Validate(valueField, ruleValue)
			if err != nil {
				validationErrors = append(validationErrors, ValidationError{Field: structFieldParam.Name, Err: err})
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func getRule(rule string) (string, string, error) {
	ruleStruct := strings.Split(rule, ":")
	if len(ruleStruct) != 2 {
		return "", "", ErrorNotValidDefinition
	}
	return ruleStruct[0], ruleStruct[1], nil
}

func getValidator(ruleName string) Validator {
	var v Validator
	switch ruleName {
	case "min":
		v = &MinValidator{}
	case "max":
		v = &MaxValidator{}
	case "len":
		v = &LenValidator{}
	case "regexp":
		v = &RegexValidator{}
	case "in":
		v = &InValidator{}
	}
	return v
}
