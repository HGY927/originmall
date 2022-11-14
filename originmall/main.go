package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"originmall/controller"
	_ "originmall/middleware/utils"
)

func main() {

	r := mux.NewRouter()
	r = r.PathPrefix("/user").Subrouter()
	r.HandleFunc("/register", controller.RegisterController).Methods("post")
	r.HandleFunc("/login", controller.LoginController).Methods("post")
	http.ListenAndServe(":8080", r)

}
