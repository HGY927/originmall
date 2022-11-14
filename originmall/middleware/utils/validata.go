package utils

import (
	"github.com/go-playground/validator/v10"
)

func CheckField(any interface{}) map[string]string {
	validate := validator.New()
	erred := validate.Struct(any)
	Map := make(map[string]string, 10)
	if erred != nil {
		for _, err := range erred.(validator.ValidationErrors) {
			if err.ActualTag() == "min" {
				Map[err.StructField()] = "输入的字符小于" + err.Param()
			} else if err.ActualTag() == "max" {
				Map[err.StructField()] = "输入的字符大于" + err.Param()
			} else if err.ActualTag() == "required" {
				Map[err.StructField()] = "字段为必填"
			}
		}
		return Map
	}
	return Map
}
