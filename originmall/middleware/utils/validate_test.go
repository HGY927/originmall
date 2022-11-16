package utils

import (
	"fmt"
	"testing"
)

type user struct {
	Username     string `json:"username" validate:"required,min=2,max=10"`
	Password     string `json:"password" validate:"required,min=6,max=20"`
	Registertime int64  `json:"registertime"`
}

func TestValidate(t *testing.T) {
	user := user{
		Username: "1",
		Password: "1",
	}

	fmt.Println(CheckField(user))

	//for k, v := range arr {
	//	fmt.Println(k, v)
	//}
}
