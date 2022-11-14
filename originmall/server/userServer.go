package server

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"originmall/dao"
	"originmall/moudle"
	"originmall/reponse"
	"time"
)

type UserServer struct {
	Username     string `json:"username" validate:"min=2,max=10,required"`
	Password     string `json:"password" validate:"min=6,max=20,required"`
	Registertime int64  `json:"registertime"`
	Logofftime   int64  `json:"logofftime"`
}

// Register 用户注册逻辑
func Register(server *UserServer) reponse.ReponseMessge {
	if dao.QueryUserByName(server.Username) {
		code := reponse.ErrorExistUser
		return reponse.ReponseMessge{
			Code:    code,
			Message: "用户名已存在",
		}
	}
	server.setRegisterTime()
	server.setHashPassword(server.Password)

	//组装dao
	user := &moudle.User{
		Username:     server.Username,
		Password:     server.Password,
		Registertime: server.Registertime,
	}
	//新增用户
	if !dao.CreateNewUser(user) {
		code := reponse.INSERTUSERERR
		return reponse.ReponseMessge{
			Code:    code,
			Message: "新增异常",
		}
	}
	code := reponse.SUCCES
	return reponse.ReponseMessge{
		Code:    code,
		Message: "新增用户成功",
	}
}

func (this *UserServer) setHashPassword(password string) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		fmt.Println("加密失败")
	}
	this.Password = string(hashPassword)
}

func (this *UserServer) setRegisterTime() {
	this.Registertime = time.Now().Unix()
}
func (this *UserServer) setLogOffTime() {
	this.Registertime = time.Now().Unix()
}
