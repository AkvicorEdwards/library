package user

import (
	"encoding/base64"
	"fmt"
	"library/config"
	"library/encryption"
	"library/mysql"
)

func FindUserByUsernameAndPassword(username string, password string) (result mysql.TableUser, exist bool) {
	k, err := encryption.AesEncryptCBC([]byte(password), []byte(config.Server.Key))
	exist = false
	if err != nil {
		fmt.Println(err)
		return
	}
	password = base64.StdEncoding.EncodeToString(k)

	db, err := mysql.Open(mysql.DEFAULT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer mysql.Close(db)

	if db.Raw("SELECT * FROM users WHERE user_name=? AND user_pwd=? LIMIT 1", username, password).Scan(&result).Error == nil {
		exist = true
	}
	return
}

func AddUser(user string, password string) bool {
	k, err := encryption.AesEncryptCBC([]byte(password), []byte(config.Server.Key))
	if err != nil {
		fmt.Println(err)
		return false
	}
	password = base64.StdEncoding.EncodeToString(k)

	db, err := mysql.Open(mysql.DEFAULT)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer mysql.Close(db)

	if db.Exec("INSERT INTO users (user_name, user_pwd) VALUES (?, ?)", user, password).Error != nil {
		db.Exec("alter table users AUTO_INCREMENT=1")
		return false
	}
	return true
}
