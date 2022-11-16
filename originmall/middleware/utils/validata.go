package utils

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func CheckField(filed any) string {
	validate := validator.New()
	erred := validate.Struct(filed)
	var str []string
	if erred != nil {
		for _, err := range erred.(validator.ValidationErrors) {
			if err.ActualTag() == "required" {
				str = append(str, err.StructField()+"字段不可为空")
			} else if err.ActualTag() == "min" {
				str = append(str, err.StructField()+"值不可小于"+err.Param())
			} else if err.ActualTag() == "max" {
				str = append(str, err.StructField()+"值不可大于"+err.Param())
			}
		}

		return strings.Join(str, ",")
	}
	return ""
}
