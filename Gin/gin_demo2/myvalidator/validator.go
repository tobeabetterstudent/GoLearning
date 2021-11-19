package myvalidator

import (
	"github.com/go-playground/validator/v10"
)

func NameValid(v validator.FieldLevel) bool {
	value,_ := v.Field().Interface().(string)
	return value != "admin"
}