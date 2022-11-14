package utils

import "github.com/go-playground/validator/v10"

func CheckField(any interface{}) error {
	validate := validator.New()
	return validate.Struct(any)
}
