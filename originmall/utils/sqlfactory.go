package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mall")
	if err = Db.Ping(); err != nil {
		fmt.Println("连接异常")
	}
	fmt.Println("sql连接正常")
}
