package server

import "testing"

func TestRegister(t *testing.T) {
	userserver := UserServer{
		Username: "李哈哈",
		Password: "123456",
	}
	Register(&userserver)
}
