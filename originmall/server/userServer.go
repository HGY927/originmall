package server

import (
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

	if ok, res := dao.QueryUserByName(server.Username); ok {
		return res
	}
	server.setRegisterTime()
	if ok, res := server.setHashPassword(server.Password); !ok {
		return res
	}
	//组装dao
	user := &moudle.User{
		Username:     server.Username,
		Password:     server.Password,
		Registertime: server.Registertime,
	}
	//新增用户
	_, res := dao.CreateNewUser(user)
	return res

}

func (this *UserServer) setHashPassword(password string) (bool, reponse.ReponseMessge) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, reponse.ReponseMessge{
			Code:    reponse.ENCODEPWDERR,
			Message: "加密异常",
		}
	}
	this.Password = string(hashPassword)
	return true, reponse.ReponseMessge{
		Code:    reponse.SUCCES,
		Message: "加密成功",
	}
}

func (this *UserServer) setRegisterTime() {
	this.Registertime = time.Now().Unix()
}
func (this *UserServer) setLogOffTime() {
	this.Registertime = time.Now().Unix()
}
