package user

import (
	"encoding/base64"
	"fmt"
	"library/config"
	"library/models/utilities/database"
	"library/models/utilities/tools"
)

func FindUserByUsernameAndPassword(username string, password string) (database.TableUser, bool) {
	k, err := tools.AesEncryptCBC([]byte(password), []byte(config.Data.Server.Key))
	if err != nil {
		fmt.Println(err)
		return database.TableUser{}, false
	}
	password = base64.StdEncoding.EncodeToString(k)
	rows, err := database.DB.Query("SELECT * FROM users WHERE user_name=? AND user_pwd=? LIMIT 1", username, password)
	if err != nil {
		fmt.Println(err)
		return database.TableUser{}, false
	}
	row := database.TableUserSQL{}
	if rows.Next() {
		if err = rows.Scan(&row.Id, &row.Username, &row.Password, &row.NickName, &row.Email, &row.Password); err != nil {
			fmt.Println(err)
			return database.TableUser{}, false
		}
	} else {
		return database.TableUser{}, false
	}
	return row.Transfer(), true
}

func AddUser(user string, password string) bool {
	k, err := tools.AesEncryptCBC([]byte(password), []byte(config.Data.Server.Key))
	if err != nil {
		fmt.Println(err)
		return false
	}
	password = base64.StdEncoding.EncodeToString(k)
	_, err = database.DB.Exec("INSERT INTO users (user_name, user_pwd) VALUES (?, ?)", user, password)
	if err != nil {
		_, err = database.DB.Exec("alter table users AUTO_INCREMENT=1")
		return false
	}
	return true
}


