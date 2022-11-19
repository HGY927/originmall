package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"originmall/controller"
	"originmall/utils"
)

var R = mux.NewRouter()

func Run() {
	//处理静态资源
	R.PathPrefix("/view/").Handler(http.StripPrefix("/view/", http.FileServer(http.Dir("template/view/"))))
	R.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("template/static/"))))
	//首页登录地址
	R.HandleFunc("/portal", func(w http.ResponseWriter, r *http.Request) {
		utils.Temp.ExecuteTemplate(w, "index.html", nil)
	})

	//路由分组，统一前缀: /user
	S := R.PathPrefix("/user").Subrouter()
	S.HandleFunc("/register", controller.RegisterController).Methods("post")
	S.HandleFunc("/login", controller.LoginController).Methods("post")
	S.HandleFunc("/update", controller.UpdateNameController).Methods("get")
}
