package dam

import "log"

// 向 user 表添加记录
func AddUser(username, nickname, password string) (int, bool) {
	if !Connected {
		Connect()
	}
	lockUser.Lock()
	defer lockUser.Unlock()

	user := User{
		Uuid:       GetInc("user") + 1,
		Username:   username,
		Nickname:   nickname,
		Password:   password,
	}

	if err := db.Table("user").Create(&user).Error; err != nil {
		log.Println(err)
		return 0, false
	}

	UpdateInc("user", user.Uuid)

	return user.Uuid, true
}

// 从 user 表中获取一条记录
func GetUser(key interface{}) *User {
	if !Connected {
		Connect()
	}
	user := &User{}
	query := ""
	switch key.(type) {
	case int, int64, int32, int16, int8:
		query = "uuid=?"
	case string:
		query = "username=?"
	default:
		return user
	}
	db.Table("user").Where(query, key).First(user)

	return user
}

