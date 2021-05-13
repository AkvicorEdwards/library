package handler

import (
	"library/session"
	"net/http"
)

var public = make(map[string]func(http.ResponseWriter, *http.Request))
var private = make(map[string]func(http.ResponseWriter, *http.Request))

func ParsePrefix() {
	// API
	public["/login"] = login
	public["/register"] = register
	private["/api/shelf"] = shelf
	private["/api/like"] = like
	private["/api/like/add"] = likeAdd
	private["/api/like/del"] = likeDel
	private["/api/like/fix"] = likeFix
	private["/api/book"] = book
	private["/api/book/add"] = bookAdd
	private["/api/book/fix"] = bookFix
	private["/api/book/like"] = bookLike
	private["/api/book/read/start"] = bookReadStart
	private["/api/book/read/finish"] = bookReadFinish
	private["/api/book/like/add"] = bookLikeAdd
	private["/api/book/like/del"] = bookLikeDel
	private["/api/book/like/fix"] = bookLikeFix

	// Page
	public["/page/login"] = loginPage
	public["/page/register"] = registerPage
	private["/page/book/add"] = pageBookAdd
	private["/page/shelf"] = pageShelf
	private["/page/book"] = pageBook
	private["/page/likes"] = pageLikes
	private["/"] = index
}

type MyHandler struct{}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := public[r.URL.Path]; ok {
		h(w, r)
		return
	}

	if !session.Verify(r) {
		public["/page/login"](w, r)
		return
	}

	if h, ok := private[r.URL.Path]; ok {
		h(w, r)
		return
	}
}
