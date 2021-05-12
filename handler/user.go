package handler

import (
	"fmt"
	"github.com/AkvicorEdwards/util"
	"library/cookie"
	"library/db"
	"library/session"
	"net/http"
	"time"
)


func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		url := r.FormValue("url")
		if len(url) == 0 || url=="/?" || url=="?" {
			url = "/"
		}
		if url[len(url)-1] == '?' {
			url = url[:len(url)-1]
		}
		if db.Login(username, password) {
			user := db.GetUserInfoByUsername(username)
			cookie.SetUserInfo(w, user)
			key := fmt.Sprintf("%x", util.SHA256String(username+password+fmt.Sprint(time.Now().UnixNano())))
			session.Update(w, r, user, key)
			Fprint(w, TplRedirect(url))
			return
		}
		Fprint(w, TplRedirect("/"))
		return
	}
	tpl := `<!DOCTYPE html>
<title>Login</title>
<form action="/page/login" method="post">
	<input type="hidden" name="url" id="url" value="%s">
	<label>Username:<input type="text" name="username" value="Akvicor"></label><br/><br/>
	<label>Password:<input type="password" name="password" value="password"></label><br/><br/>
	<input type="submit" value="Login">
</form>`

	Fprintf(w, tpl, fmt.Sprintf("%s?%s", r.URL.Path, r.URL.RawQuery))
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if db.Login(username, password) {
		user := db.GetUserInfoByUsername(username)
		cookie.SetUserInfo(w, user)
		key := fmt.Sprintf("%x", util.SHA256String(username+password+fmt.Sprint(time.Now().UnixNano())))
		session.Update(w, r, user, key)
		Fprintf(w, `{"status":"%d", "key":"%s"}`, StatusOK, key)
	} else {
		Fprintf(w, ResponseER)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	nickname := r.FormValue("nickname")
	password := r.FormValue("password")
	profile := r.FormValue("profile_photo")
	err := db.AddUser(username, password, profile, nickname)
	if err != nil {
		Fprintf(w, ResponseER)
		Fprintf(w, ResponseER)
		return
	}
	Fprintf(w, ResponseOK)
}
