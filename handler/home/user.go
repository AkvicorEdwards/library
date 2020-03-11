package home

import (
	"fmt"
	"html/template"
	"net/http"
	"library/config"
	user2 "library/models/user"
	"library/session"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles(config.Data.Path.Theme + "login.tpl")

		phrase := map[string]interface{}{
			"lang": "zh-cn",
			"title": config.Data.Server.Title,
		}

		if err := t.Execute(w, phrase); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, ok := user2.FindUserByUsernameAndPassword(username, password)
	if !ok {
		_, _ = fmt.Fprintln(w, "User do not exist")
		return
	}
	sess := session.GetSession(w, r)
	sess.SetAttr("user", user)
	http.Redirect(w, r, "/", 302)
}
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles(config.Data.Path.Theme + "register.tpl")

		phrase := map[string]interface{}{
			"lang": "zh-cn",
			"title": config.Data.Server.Title,
		}

		if err := t.Execute(w, phrase); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	ok := user2.AddUser(username, password)
	if !ok {
		_, _ = fmt.Fprintln(w, "ERROR")
		return
	}

	http.Redirect(w, r, "/", 302)
}
