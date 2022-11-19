package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
	"originmall/reponse"
	"originmall/server"
	"originmall/utils"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("SESSION"); err == nil {
		fmt.Println("已经登录，直接跳转至首页", cookie)
		utils.Temp.ExecuteTemplate(w, "index.html", nil)
	}
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
	reps := server.Login(userserver)
	if reps.Code == reponse.SUCCES {
		//设置seesion
		Store := sessions.NewCookieStore(securecookie.GenerateRandomKey(32))
		session, err := Store.Get(r, "SESSION")
		if err != nil {
			fmt.Println("用户session生成失败", err.Error())
			return
		}
		session.ID = utils.UuidValue()
		session.Values["name"] = userserver.Username
		session.Values["time"] = reps.Data
		session.Options.HttpOnly = true
		session.Options.MaxAge = 86400
		session.Options.Secure = true
		err = Store.Save(r, w, session)
		if err != nil {
			fmt.Println("保存用户session异常")
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reps)
	utils.Temp.ExecuteTemplate(w, "index.html", nil)
}
