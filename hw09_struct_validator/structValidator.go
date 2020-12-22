package hw09_struct_validator //nolint:golint,stylecheck

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Validator interface {
	Validate(valueField reflect.Value, ruleValue string) error
}

type MinValidator struct {
	min  int
	cond int
}

type MaxValidator struct {
	max  int
	cond int
}

type LenValidator struct {
	len  int
	cond string
}

type RegexValidator struct {
	cond string
}

type InValidator struct {
	cond interface{}
}

func (vIn *InValidator) Validate(valueField reflect.Value, ruleValue string) error {
	inValues := strings.Split(ruleValue, ",")
	switch valueField.Kind() { // nolint: exhaustive
	case reflect.String:
		cond := valueField.String()
		for _, inValue := range inValues {
			if inValue == cond {
				return nil
			}
		}
		return &InValidator{cond}
	case reflect.Int:
		cond := int(valueField.Int())
		for _, inValue := range inValues {
			intInValue, _ := strconv.Atoi(inValue)
			if intInValue == cond {
				return nil
			}
		}
		return &InValidator{cond}
	default:
		return nil
	}
}

func (vReg *RegexValidator) Validate(valueField reflect.Value, ruleValue string) error {
	emailRegEx := regexp.MustCompile(ruleValue)
	cond := valueField.String()

	resReg := emailRegEx.MatchString(cond)
	if resReg {
		return nil
	}

	return &RegexValidator{cond}
}

func (vMax *MaxValidator) Validate(valueField reflect.Value, ruleValue string) error {
	max, err := strconv.Atoi(ruleValue)
	if err != nil {
		return err
	}

	cond := int(valueField.Int())

	if cond <= max {
		return nil
	}

	return &MaxValidator{max, cond}
}

func (vMin *MinValidator) Validate(valueField reflect.Value, ruleValue string) error {
	min, err := strconv.Atoi(ruleValue)
	if err != nil {
		return err
	}

	cond := int(valueField.Int())

	if cond >= min {
		return nil
	}

	return &MinValidator{min, cond}
}

func (vLen *LenValidator) Validate(valueField reflect.Value, ruleValue string) error {
	lenString, err := strconv.Atoi(ruleValue)
	if err != nil {
		return err
	}

	switch valueField.Kind() { // nolint: exhaustive
	case reflect.String:
		cond := valueField.String()
		if len(cond) == lenString {
			return nil
		}
		return &LenValidator{lenString, cond}
	case reflect.Slice:
		for i := 0; i < valueField.Len(); i++ {
			sliceValue := valueField.Index(i)
			cond := sliceValue.String()
			if len(cond) != lenString {
				return &LenValidator{lenString, cond}
			}
		}
		return nil
	}

	return nil
}

func (vMin *MinValidator) Error() string {
	return fmt.Sprintf("[min] %d <= %d", vMin.cond, vMin.min)
}

func (vMax *MaxValidator) Error() string {
	return fmt.Sprintf("[max] %d >= %d", vMax.cond, vMax.max)
}

func (vLen *LenValidator) Error() string {
	return fmt.Sprintf("[len] values longer then %s != %d", vLen.cond, vLen.len)
}

func (vReg *RegexValidator) Error() string {
	return fmt.Sprintf("[regex] not valid email %s", vReg.cond)
}

func (vIn *InValidator) Error() string {
	return fmt.Sprintf("[in] value %s not in", vIn.cond)
}
