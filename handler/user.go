package handler

import (
	"fmt"
	"github.com/AkvicorEdwards/util"
	"library/cookie"
	"library/db"
	"library/def"
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
	<label>Username:<input type="text" name="username"></label><br/><br/>
	<label>Password:<input type="password" name="password"></label><br/><br/>
	<input type="submit" value="Login">
</form>`

	Fprintf(w, tpl, fmt.Sprintf("%s?%s", r.URL.Path, r.URL.RawQuery))
	return
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	if !def.PermitRegister {
		return
	}
	if r.Method == "POST" {
		username := r.FormValue("username")
		if len(username) <= 0 {
			return
		}
		password := r.FormValue("password")
		if len(password) <= 0 {
			return
		}
		nickname := r.FormValue("nickname")
		if len(nickname) <= 0 {
			return
		}
		profile := r.FormValue("profile")
		url := r.FormValue("url")
		if len(url) == 0 || url=="/?" || url=="?" {
			url = "/"
		}
		if url[len(url)-1] == '?' {
			url = url[:len(url)-1]
		}

		if db.AddUser(username,password,profile,nickname) == nil {
			Fprint(w, TplRedirect("/"))
			return
		}
		Fprint(w, TplRedirect(url))
		return
	}
	tpl := `<!DOCTYPE html>
<title>Register</title>
<form action="/page/register" method="post">
	<input type="hidden" name="url" id="url" value="%s">
	<label>Username:<input type="text" name="username" required></label><br/><br/>
	<label>Password:<input type="password" name="password" required></label><br/><br/>
	<label>Nickname:<input type="text" name="nickname" required></label><br/><br/>
	<label>Profile Photo:<input type="text" name="profile"></label><br/><br/>
	<input type="submit" value="Register">
</form>`

	Fprintf(w, tpl, fmt.Sprintf("%s?%s", r.URL.Path, r.URL.RawQuery))
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	if len(username) <= 0 {
		Fprintf(w, ResponseER)
		return
	}
	password := r.FormValue("password")
	if len(password) <= 0 {
		Fprintf(w, ResponseER)
		return
	}
	if db.Login(username, password) {
		user := db.GetUserInfoByUsername(username)
		cookie.SetUserInfo(w, user)
		key := fmt.Sprintf("%x", util.SHA256String(username+password+fmt.Sprint(time.Now().UnixNano())))
		session.Update(w, r, user, key)
		Fprint(w, MarshalLoginResponseInfo(key, user))
	} else {
		Fprint(w, ResponseER)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	if !def.PermitRegister {
		Fprintf(w, ResponseER)
		return
	}
	username := r.FormValue("username")
	if len(username) <= 0 {
		Fprintf(w, ResponseER)
		return
	}
	nickname := r.FormValue("nickname")
	if len(nickname) <= 0 {
		Fprintf(w, ResponseER)
		return
	}
	password := r.FormValue("password")
	if len(password) <= 0 {
		Fprintf(w, ResponseER)
		return
	}
	profile := r.FormValue("profile_photo")
	err := db.AddUser(username, password, profile, nickname)
	if err != nil {
		Fprintf(w, ResponseER)
		return
	}
	Fprintf(w, ResponseOK)
}
