package handler

import (
	"fmt"
	"library/config"
	"library/handler/book"
	"library/handler/home"
	"library/session"
	"net/http"
	"regexp"
)

type str2func map[string]func(http.ResponseWriter, *http.Request)

var mux str2func
var api str2func

func ParsePrefix() {
	mux = make(str2func)
	api = make(str2func)

	// User action
	api["/login"] = home.Login 		// login.tpl
	api["/register"] = home.Register 	// register.tpl

	// Home Page
	mux["/"] = home.Page	 // home.tpl

	// Book
	mux["/add/book"] = book.AddBook
	mux["/add/favour"] = book.AddFavour
	mux["/set/time"] = book.SetTime
	mux["/set/start/read"] = book.SetStartRead
	mux["/fix"] = book.Fix
	mux["/fix/cover"] = book.FixCover
	mux["/books"] = book.Index
	mux["/book"] = book.View

}

type MyHandler struct {}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := api[r.URL.Path]; ok {
		h(w, r)
		return
	}
	if _, ok := mux[r.URL.Path]; !ok {
		match := func(pattern string) (matched bool) {
			matched, _ =  regexp.MatchString(pattern, r.URL.String())
			return
		}

		fileServer := func(prefix string, dir string) {
			http.StripPrefix(prefix, http.FileServer(http.Dir(config.Data.Path.Theme + dir))).ServeHTTP(w, r)
		}

		if match("/css/") {
			fileServer("/css/", "css/")
		} else if match("/js/") {
			fileServer("/js/", "js/")
		} else if match("/img/") {
			fileServer("/img/", "img/")
		} else if match("/cover/") {
			http.StripPrefix("/cover/", http.FileServer(http.Dir(config.Data.Path.Cover))).ServeHTTP(w, r)
		} else if match("/favicon.ico") {
			fileServer("/", "img/")
		} else if match("/book/") {
			mux["/book"](w, r)
		} else {
			//fmt.Println(r.URL.Path)
			mux["/"](w, r)
		}
		return
	}
	_, exist := session.GetSession(w, r).GetAttr("user")
	if !exist {
		fmt.Println(r.URL.Path)
		home.Login(w, r)
		return
	}
	mux[r.URL.Path](w, r)
}



