package controller

import (
	"encoding/json"
	"net/http"
	"originmall/reponse"
	"originmall/server"
	"originmall/utils"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	userserver := &server.UserServer{}
	userserver.Username = r.FormValue("username")
	userserver.Password = r.FormValue("password")
	str := utils.CheckField(userserver)
	if str != "" {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(reponse.ReponseMessge{
			Code:    reponse.FILEDCHECKERR,
			Message: str,
		})
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(server.Register(userserver))
}
