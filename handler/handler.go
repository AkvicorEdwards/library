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

var public str2func
var protected str2func
var private str2func
var api str2func

func ParsePrefix() {
	public = make(str2func)
	protected = make(str2func)
	private = make(str2func)
	api = make(str2func)

	private["/"] = home.Home
	public["/login"] = home.Login
	public["/register"] = home.Register

	// Book
	private["/add/book"] = book.AddBook
	private["/add/favour"] = book.AddFavour
	private["/set/time"] = book.SetTime
	private["/set/start/read"] = book.SetStartRead
	private["/fix"] = book.Fix
	private["/fix/cover"] = book.FixCover
	private["/books"] = book.Index
	private["/book"] = book.View

}

type MyHandler struct {}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Public
	if h, ok := public[r.URL.Path]; ok {
		h(w, r)
		return
	}
	// API
	if h, ok := api[r.URL.Path]; ok {
		h(w, r)
		return
	}
	// Some content requires login to display
	if h, ok := protected[r.URL.Path]; ok {
		h(w, r)
		return
	}
	// Need to be logged in
	if h, ok := private[r.URL.Path]; ok {
		// Get session
		ses, err := session.Get(r, "user")
		if err != nil {
			_, _ = fmt.Fprintln(w, "Error 1")
			return
		}
		// Check permission
		per, ok := ses.Values["permission"].(int64)
		if !ok {
			public["/login"](w, r)
			return
		}

		if per&1 == 0 {
			_, _ = fmt.Fprintln(w, "Do not have permission")
			return
		}

		h(w, r)
		return
	}

	match := func(pattern string) (matched bool) {
		matched, _ = regexp.MatchString(pattern, r.URL.String())
		return
	}

	fileServer := func(prefix string, dir string) {
		http.StripPrefix(prefix, http.FileServer(http.Dir(config.Path.Theme+dir))).ServeHTTP(w, r)
	}

	if match("/css/") {
		fileServer("/css/", "css/")
	} else if match("/js/") {
		fileServer("/js/", "js/")
	} else if match("/cover/") {
		http.StripPrefix("/cover/", http.FileServer(http.Dir(config.Path.Cover))).ServeHTTP(w, r)
	}  else if match("/img/") {
		fileServer("/img/", "img/")
	} else if match("favicon.ico"){
		fileServer("/", "img/")
	} else if match("/book/") {
		// Get session
		ses, err := session.Get(r, "user")
		if err != nil {
			_, _ = fmt.Fprintln(w, "Error 1")
			return
		}
		// Check permission
		per, ok := ses.Values["permission"].(int64)
		if !ok {
			public["/login"](w, r)
			return
		}

		if per&1 == 0 {
			_, _ = fmt.Fprintln(w, "Do not have permission")
			return
		}
		private["/book"](w, r)
	}  else {
		fmt.Println(r.URL.Path)
	}

}




