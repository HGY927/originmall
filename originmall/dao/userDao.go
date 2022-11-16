package dao

import (
	"originmall/middleware/utils"
	"originmall/moudle"
)

var (
	namesql   = "select username from user where username= ?"
	insertsql = "insert into user (username,pwd,registertime)values(?,?,?)"
)

// QueryUserByName 查询数据库中是否存在该用户名
func QueryUserByName(username string) bool {
	var temp string
	row := utils.Db.QueryRow(namesql, username)
	row.Scan(&temp)
	if temp == "" {
		return false
	}
	return true
}

// CreateNewUser 新增用户
func CreateNewUser(user *moudle.User) bool {

	exact, _ := utils.Db.Exec(insertsql, user.Username, user.Password, user.Registertime)
	isAdd, err := exact.RowsAffected()
	if err != nil || isAdd == 0 {
		return false
	}
	return true
}
