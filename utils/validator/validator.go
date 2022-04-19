package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	validator "github.com/go-playground/validator/v10"
)

var v *validator.Validate

func Init() {
	v = validator.New()
}
func Validate(obj interface{}) error {
	err := v.Struct(obj)
	if err == nil {
		return nil
	}
	structName := getType(obj)
	errs := err.(validator.ValidationErrors)
	message := strings.ReplaceAll(errs[0].Namespace(), structName+".", "") + " is invalid or missing"
	return errors.New(fmt.Sprintf("message:%s", message))
}

func ValidateVariable(obj interface{}, tags, parameterName string) error {
	err := v.Var(obj, tags)
	if err == nil {
		return nil
	}
	message := parameterName + " is invalid or missing"
	return errors.New(fmt.Sprintf("message:%s", message))
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

/***ValidatePasswordString - validates input password string bases on the following constraints***/
func validatePasswordString(str string) bool {
	var (
		hasMinLen  = false
		hasMaxLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(str) >= 8 {
		hasMinLen = true
	}
	if len(str) <= 48 {
		hasMaxLen = true
	}
	for _, ch := range str {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsNumber(ch):
			hasNumber = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}
	return hasMinLen && hasMaxLen && hasUpper && hasLower && hasNumber && hasSpecial
}
