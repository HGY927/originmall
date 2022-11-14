package dao

import (
	_ "originmall/middleware/utils"
	"testing"
)

func TestUser(t *testing.T) {

	QueryUserByName("哈哈")

}
