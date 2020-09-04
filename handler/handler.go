package handler

import (
	"fmt"
	"io"
	"library/session"
	"net/http"
	"os"
	"regexp"
)

type str2func map[string]func(http.ResponseWriter, *http.Request)

var public str2func

func ParsePrefix() {
	public = make(str2func)

	// 界面
	public["/"] = index
	public["/login"] = login
	public["/register"] = register
	public["/add/book"] = AddBook
	public["/books"] = bookIndex
	public["/book"] = bookView

	public["/add/favour"] = bookAddFavour
	public["/set/time"] = bookSetTime
	public["/set/start/read"] = bookSetStartRead
	public["/fix"] = bookFix
	public["/fix/cover"] = bookFixCover
	public["/fix/favour"] = bookFixFavour
	public["/del/favour"] = bookDelFavour

}

type MyHandler struct{}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/register" {
		public["/register"](w, r)
		return
	}

	if session.GetUUID(r) <= 0 {
		public["/login"](w, r)
		return
	}

	if h, ok := public[r.URL.Path]; ok {
		h(w, r)
		return
	}

	match := func(pattern string) (matched bool) {
		matched, _ = regexp.MatchString(pattern, r.URL.String())
		return
	}

	if match("/favicon.ico") {
		download(w, "./favicon.ico")
	} else if match("/cover/") {
		http.StripPrefix("/cover/", http.FileServer(http.Dir("./cover/"))).ServeHTTP(w, r)
	} else if match("/book/") {
		public["/book"](w, r)
	}

}

func download(w http.ResponseWriter, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		_, _ = fmt.Fprintln(w, "File Not Found")
		return
	}
	defer func() { _ = file.Close() }()
	data := make([]byte, 1024)
	for {
		n, err1 := file.Read(data)
		if err1 != nil && err1 != io.EOF {
			_, _ = fmt.Fprintln(w, "File Read Error")
			return
		}
		nn, err2 := w.Write(data[:n])
		if err2 != nil || nn != n {
			_, _ = fmt.Fprintln(w, "File Write Error")
			return
		}
		if err1 == io.EOF {
			return
		}
	}
}

func Fprint(w http.ResponseWriter, a ...interface{}) {
	_, _ = fmt.Fprint(w, a...)
}

func Fprintf(w http.ResponseWriter, format string, a ...interface{}) {
	_, _ = fmt.Fprintf(w, format, a...)
}

func Fprintln(w http.ResponseWriter, a ...interface{}) {
	_, _ = fmt.Fprintln(w, a...)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

