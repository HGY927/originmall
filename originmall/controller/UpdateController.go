package controller

import (
	"fmt"
	"net/http"
)

func UpdateNameController(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("SESSION")
	fmt.Println(cookie.String())
}
