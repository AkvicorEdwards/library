package home

import (
	"fmt"
	"library/config"
	"library/session"
	"library/tpl"
	user2 "library/user"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	phrase := map[string]interface{}{
		"lang":  "zh-cn",
		"title": config.Server.Title,
	}

	if err := tpl.Home.Execute(w, phrase); err != nil {
		fmt.Println(err)
		_, _ = fmt.Fprintf(w, "%v", "Error")
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		phrase := map[string]interface{}{
			"lang":  "zh-cn",
			"title": config.Server.Title,
		}

		if err := tpl.Login.Execute(w, phrase); err != nil {
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

	ses, _ := session.Get(r, "user")
	ses.Values["username"] = user.UserName
	ses.Values["nickname"] = user.NickName
	ses.Values["email"] = user.Email
	ses.Values["id"] = user.Id
	ses.Values["permission"] = user.Permissions
	ses.Options.MaxAge = 60*60*24
	err := ses.Save(r, w)
	if err != nil {
		_, _ = fmt.Fprintln(w, "ERROR 2")
		return
	}
	http.Redirect(w, r, "/", 302)

}
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		phrase := map[string]interface{}{
			"lang":  "zh-cn",
			"title": config.Server.Title,
		}

		if err := tpl.Register.Execute(w, phrase); err != nil {
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
