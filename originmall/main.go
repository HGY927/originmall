package main

import (
	"net/http"
	"originmall/controller"
	_ "originmall/middleware/utils"
)

func main() {

	http.HandleFunc("/user/register", controller.UserController)
	http.ListenAndServe(":8080", nil)

}
