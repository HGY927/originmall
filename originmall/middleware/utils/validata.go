package utils

import "github.com/go-playground/validator/v10"

func CheckField(any interface{}) {
	validate := validator.New()
	validate.Struct(any)
}
