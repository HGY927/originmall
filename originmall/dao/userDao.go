package dao

import (
	"originmall/middleware/utils"
	"originmall/moudle"
	"originmall/reponse"
	"time"
)

var (
	namesql   = "select username from user where username= ?"
	insertsql = "insert into user (username,pwd,registertime)values(?,?,?)"
)

// QueryUserByName 查询数据库中是否存在该用户名
func QueryUserByName(username string) (bool, reponse.ReponseMessge) {
	var temp string
	row := utils.Db.QueryRow(namesql, username)
	row.Scan(&temp)
	if temp == "" {
		return false, reponse.ReponseMessge{
			Code:    reponse.INSERTUSER,
			Message: "可以注册",
		}
	}
	return true, reponse.ReponseMessge{
		Code:    reponse.REPEATUSER,
		Message: "重复的用户名",
	}
}

// CreateNewUser 新增用户
func CreateNewUser(user *moudle.User) (bool, reponse.ReponseMessge) {

	exact, _ := utils.Db.Exec(insertsql, user.Username, user.Password, user.Registertime)
	isAdd, err := exact.RowsAffected()
	if err != nil || isAdd == 0 {
		return false, reponse.ReponseMessge{
			Code:    reponse.INSERTUSERERR,
			Message: "新增用户异常失败",
		}
	}
	return true, reponse.ReponseMessge{
		Code:    reponse.SUCCES,
		Message: "新增用户成功",
		Data:    time.Unix(user.Registertime, 0).Format("2006-01-02 15:04:05"),
	}
}
