package server

import (
	"golang.org/x/crypto/bcrypt"
	"originmall/dao"
	"originmall/moudle"
	"originmall/reponse"
	"time"
)

type UserServer struct {
	Username     string `json:"username" validate:"required,min=2,max=10"`
	Password     string `json:"password" validate:"required,min=6,max=12"`
	Registertime int64  `json:"registertime"`
	Logofftime   int64  `json:"logofftime"`
}

// Register 用户注册逻辑
func Register(server *UserServer) reponse.ReponseMessge {

	if ok, _ := dao.QueryUserByName(server.Username); ok {
		return reponse.ReponseMessge{
			Code:    reponse.REPEATUSER,
			Message: "重复的用户名",
		}
	}
	if ok := server.setHashPassword(server.Password); !ok {
		return reponse.ReponseMessge{
			Code:    reponse.ENCODEPWDERR,
			Message: "加密异常",
		}
	}
	server.setRegisterTime()
	//组装dao
	user := &moudle.User{
		Username:     server.Username,
		Password:     server.Password,
		Registertime: server.Registertime,
	}
	//新增用户
	flag := dao.CreateNewUser(user)
	if !flag {
		return reponse.ReponseMessge{
			Code:    reponse.INSERTUSERERR,
			Message: "新增用户异常",
		}
	}
	return reponse.ReponseMessge{
		Code:    reponse.SUCCES,
		Message: "新增用户成功",
	}
}

// Login 登录用户逻辑
func Login(server *UserServer) reponse.ReponseMessge {

	ok, user := dao.QueryUserByName(server.Username)
	if !ok || !server.checkUserPassword(user, server.Password) {
		return reponse.ReponseMessge{
			Code:    reponse.NOTCORRECTUSERNAMEORPASSWORD,
			Message: "用户名或者密码不正确",
		}
	}
	return reponse.ReponseMessge{
		Code:    reponse.SUCCES,
		Message: "登录成功",
	}

}

// 密码加密
func (this *UserServer) setHashPassword(password string) bool {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	this.Password = string(hashPassword)
	return err == nil
}

// 校验密码
func (this *UserServer) checkUserPassword(user *moudle.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (this *UserServer) setRegisterTime() {
	this.Registertime = time.Now().Unix()
}
func (this *UserServer) setLogOffTime() {
	this.Registertime = time.Now().Unix()
}
