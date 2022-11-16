package dao

import (
	"originmall/moudle"
	"originmall/utils"
)

var (
	namesql   = "select username,pwd password,registertime from user where username= ?"
	insertsql = "insert into user (username,pwd,registertime)values(?,?,?)"
)

// QueryUserByName 查询数据库中是否存在该用户名
func QueryUserByName(username string) (bool, *moudle.User) {
	row := utils.Db.QueryRow(namesql, username)
	user := &moudle.User{}
	row.Scan(&user.Username, &user.Password, &user.Registertime)
	return user.Username != "", user
}

// CreateNewUser 新增用户
func CreateNewUser(user *moudle.User) bool {

	exact, _ := utils.Db.Exec(insertsql, user.Username, user.Password, user.Registertime)
	isAdd, err := exact.RowsAffected()
	return !(err != nil || isAdd == 0)
}
