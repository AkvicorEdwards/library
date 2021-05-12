package session

import (
	"fmt"
	"github.com/gorilla/sessions"
	"library/db"
	"library/def"
	"net/http"
)

func GetUUID(r *http.Request) int64 {
	ses, err := Get(r, def.SessionName)
	if err != nil {
		return 0
	}
	uuid, ok := ses.Values["uuid"].(int64)
	if !ok {
		return 0
	}
	return uuid
}

func Update(w http.ResponseWriter, r *http.Request, user *db.User, key string) {
	ses, _ := GetUser(r)
	ses.Values["uuid"] = user.UUID
	ses.Values["username"] = user.Username
	ses.Values["nickname"] = user.Nickname
	if len(key) != 0 {
		ses.Values["key"] = key
	}

	ses.Options.MaxAge = 60 * 60 * 24
	err := ses.Save(r, w)
	if err != nil {
		_, _ = fmt.Fprintln(w, "ERROR session SetUserInfo")
		return
	}
}

func Verify(r *http.Request) bool {
	ses, err := Get(r, def.SessionName)
	if err != nil {
		return false
	}
	key, ok := ses.Values["key"].(string)
	if !ok {
		return false
	}
	return len(key) != 0
}

func GetUser(r *http.Request) (*sessions.Session, error) {
	return Get(r, def.SessionName)
}
