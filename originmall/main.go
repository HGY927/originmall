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
	r.HandleFunc("/register", controller.UserController).Methods("post")
	http.ListenAndServe(":8080", r)

}
