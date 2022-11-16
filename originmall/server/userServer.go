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

	if ok := dao.QueryUserByName(server.Username); ok {
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

func Login(server *UserServer) reponse.ReponseMessge {
	return reponse.ReponseMessge{}

}

// 密码加密
func (this *UserServer) setHashPassword(password string) bool {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	this.Password = string(hashPassword)
	return true
}

func (this *UserServer) setRegisterTime() {
	this.Registertime = time.Now().Unix()
}
func (this *UserServer) setLogOffTime() {
	this.Registertime = time.Now().Unix()
}
