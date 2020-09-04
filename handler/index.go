package handler

import (
	"fmt"
	"library/cookie"
	"library/def"
	"library/operator"
	"library/session"
	"library/tpl"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.Index.Execute(w, nil); err != nil {
			fmt.Println(err)
			Fprintf(w, "%v", "Error")
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.Login.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, ok := operator.Login(username, password)

	if !ok {
		_, _ = fmt.Fprintln(w, "用户不存在或密码错误")
		return
	}

	session.SetUserInfo(w, r, user)
	cookie.SetUserInfo(w, user)

	http.Redirect(w, r, "/", 302)

}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.Register.Execute(w, nil); err != nil {
			fmt.Println(err)
			Fprintf(w, "%v", "Error")
		}
		return
	}

	username := r.FormValue("username")
	nickname := r.FormValue("nickname")
	password := r.FormValue("password")
	permissionCode := r.FormValue("p_code")
	if permissionCode != def.RegisterCode {
		Fprintln(w, "ERROR")
		return
	}

	ok := operator.AddUser(username, nickname, password)
	if !ok {
		Fprintln(w, "ERROR")
		return
	}

	http.Redirect(w, r, "/", 302)
}
