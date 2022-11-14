package controller

import (
	"encoding/json"
	"net/http"
	"originmall/middleware/utils"
	"originmall/reponse"
	"originmall/server"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	var userserver server.UserServer
	json.NewDecoder(r.Body).Decode(&userserver)
	arr := utils.CheckField(&userserver)
	if len(arr) != 0 {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(reponse.ReponseMessge{
			Code:    reponse.FILEDCHECKERR,
			Message: "字段校验失败",
			Data:    json.Marshal(arr),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(server.Register(&userserver))

}
