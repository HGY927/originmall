package main

import (
	"net/http"
	"originmall/router"
)

func main() {
	router.Run()
	http.ListenAndServe(":8080", router.R)

}
