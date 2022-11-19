package utils

import "html/template"

var Temp *template.Template

func init() {
	Temp = template.Must(template.ParseGlob("template/view/*.html"))
}
