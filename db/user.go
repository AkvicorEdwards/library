package db

import "time"

const TableUser = "user"

type User struct {
	UUID int64 `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	ProfilePhoto string `json:"profile_photo"`
	Nickname string `json:"nickname"`
	LastLogin int64 `json:"last_login"`
}

func GetUserInfo(uuid int64) *User {
	if !Connected {
		Connect()
	}
	lockUser.RLock()
	defer lockUser.RUnlock()
	user := &User{}
	err := db.Table(TableUser).Where("uuid=?", uuid).First(user).Error
	if err != nil {
		return nil
	}
	return user
}

func GetUserInfoByUsername(username string) *User {
	if !Connected {
		Connect()
	}
	lockUser.RLock()
	defer lockUser.RUnlock()
	user := &User{}
	err := db.Table(TableUser).Where("username=?", username).First(user).Error
	if err != nil {
		return nil
	}
	return user
}

func AddUser(username, password, profile, nickname string) error {
	if !Connected {
		Connect()
	}
	lockUser.Lock()
	defer lockUser.Unlock()
	user := &User{
		UUID:         GetInc(TableUser) + 1,
		Username:     username,
		Password:     password,
		ProfilePhoto: profile,
		Nickname:     nickname,
	}
	if err := db.Table(TableUser).Create(user).Error; err != nil {
		return err
	}
	updateInc(TableUser, user.UUID)
	_ = createBookshelf(user.UUID)
	_ = createLike(user.UUID)
	_ = createBookRead(user.UUID)
	_ = createBookLike(user.UUID)

	return nil
}

func Login(username, password string) bool {
	if !Connected {
		Connect()
	}
	lockUser.RLock()
	defer lockUser.RUnlock()
	res := db.Table(TableUser).Where("username=? AND password=?", username, password).Update("last_login", time.Now().Unix())
	if res.Error != nil {
		return false
	}
	if res.RowsAffected != 1 {
		return false
	}
	return true
}
