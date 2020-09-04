package operator

import (
	"library/dam"
	"library/def"
)

// 添加一个用户
func AddUser(username, nickname, password string) bool {

	if !(def.CheckUsername(username) && def.CheckPassword(password)) {
		return false
	}

	if _, ok := dam.AddUser(username, nickname, password); !ok {
		return false
	}

	return true
}

// 登录
func Login(username, password string) (*dam.User, bool) {
	user := dam.GetUser(username)
	if user.Uuid <= 0 || user.Password != password {
		return &dam.User{}, false
	}
	return user, true
}

// 获取一条用户记录
func GetUser(uuid int) (*dam.User, bool) {
	user := dam.GetUser(uuid)
	if user.Uuid <= 0 {
		return &dam.User{}, false
	}
	return user, true
}
