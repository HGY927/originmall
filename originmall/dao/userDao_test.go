package dao

import (
	"fmt"
	_ "originmall/utils"
	"testing"
)

func TestUser(t *testing.T) {

	flag, ptr := QueryUserByName("何国洋")
	fmt.Println(flag)
	fmt.Println(*ptr)

}
