package controller

import (
	"encoding/json"
	"net/http"
	"originmall/server"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	var userserver server.UserServer
	json.NewDecoder(r.Body).Decode(&userserver)

}
