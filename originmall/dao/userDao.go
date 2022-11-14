package dao

import (
	"originmall/middleware/utils"
	"originmall/moudle"
)

var (
	namesql   = "select username,pwd,registertime from user where username= ?"
	insertsql = "insert into user (username,pwd,registertime)values(?,?,?)"
)

// QueryUserByName 查询数据库中是否存在该用户名
func QueryUserByName(username string) bool {

	row := utils.Db.QueryRow(namesql, username)
	user := moudle.User{}
	row.Scan(&user.Username, &user.Password, &user.Registertime)
	if user.Username == "" {
		//fmt.Println("不存在用户名，可以注册")
		return false
	}
	return true
}

// CreateNewUser 新增用户
func CreateNewUser(user *moudle.User) bool {

	exact, _ := utils.Db.Exec(insertsql, user.Username, user.Password, user.Registertime)
	isAdd, err := exact.RowsAffected()
	if err != nil || isAdd == 0 {
		//fmt.Println("新增用户异常失败", err)
		return false
	}
	return true
}
