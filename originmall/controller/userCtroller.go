package controller

import (
	"encoding/json"
	"net/http"
	"originmall/server"
)

func UserController(w http.ResponseWriter, r *http.Request) {

	userserver := &server.UserServer{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	js, err := json.Marshal(server.Register(userserver))
	if err != nil {
		w.Write([]byte("解析数据异常"))
	}
	w.Write(js)

}
