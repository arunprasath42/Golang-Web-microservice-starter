package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

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
