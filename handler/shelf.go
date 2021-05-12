package handler

import (
	"library/db"
	"library/session"
	"net/http"
	"strings"
)

func shelf(w http.ResponseWriter, r *http.Request) {
	uuid := session.GetUUID(r)
	if uuid <= 0 {
		return
	}
	bookshelf := db.GetBookList(uuid)
	Fprint(w, Marshal(bookshelf))
}

func like(w http.ResponseWriter, r *http.Request) {
	uuid := session.GetUUID(r)
	if uuid <= 0 {
		return
	}
	likes := db.GetLikes(uuid)
	Fprint(w, Marshal(likes))
}

func likeAdd(w http.ResponseWriter, r *http.Request) {
	uuid := session.GetUUID(r)
	if uuid <= 0 {
		return
	}
	origin := strings.TrimSpace(r.FormValue("origin"))
	content := strings.TrimSpace(r.FormValue("content"))
	comment := strings.TrimSpace(r.FormValue("comment"))
	if db.AddLike(uuid, origin, content, comment) != nil {
		Fprint(w, ResponseER)
		return
	}
	Fprint(w, ResponseOK)
}

func likeDel(w http.ResponseWriter, r *http.Request) {
	uuid := session.GetUUID(r)
	if uuid <= 0 {
		return
	}
	id := Int64(strings.TrimSpace(r.FormValue("id")))
	if db.DeleteLike(uuid, id) != nil {
		Fprint(w, ResponseER)
		return
	}
	Fprint(w, ResponseOK)
}

func likeFix(w http.ResponseWriter, r *http.Request) {
	uuid := session.GetUUID(r)
	if uuid <= 0 {
		return
	}
	id := Int64(strings.TrimSpace(r.FormValue("id")))
	origin := strings.TrimSpace(r.FormValue("origin"))
	content := strings.TrimSpace(r.FormValue("content"))
	comment := strings.TrimSpace(r.FormValue("comment"))
	if db.UpdateLike(uuid, id, origin, content, comment) != nil {
		Fprint(w, ResponseER)
		return
	}
	Fprint(w, ResponseOK)
}

func book(w http.ResponseWriter, r *http.Request) {
	uuid := session.GetUUID(r)
	if uuid <= 0 {
		return
	}
	bid := Int64(r.FormValue("b"))
	bookInf := db.GetBook(uuid, bid)
	Fprint(w, Marshal(bookInf))
}

func bookFix(w http.ResponseWriter, r *http.Request) {
	uuid := session.GetUUID(r)
	if uuid <= 0 {
		return
	}
	bid := Int64(r.FormValue("b"))
	book := db.Book{
		Id:          bid,
		Title:       r.FormValue("title"),
		TitleOrigin: r.FormValue("title-origin"),
		Author:      FilterStringBySep(r.FormValue("author"), ","),
		Translator:  FilterStringBySep(r.FormValue("translator"), ","),
		Publisher:   r.FormValue("publisher"),
		Cover:       r.FormValue("cover"),
		Tag:         FilterStringBySep(r.FormValue("tag"), ","),
		Reads:       nil,
		Likes:       nil,
	}
	err := db.UpdateBook(uuid, book)
	if err != nil {
		Fprint(w, ResponseER)
		return
	}
	Fprint(w, ResponseOK)
}

func bookAdd(w http.ResponseWriter, r *http.Request) {
	uuid := session.GetUUID(r)
	if uuid <= 0 {
		return
	}
	book := db.Book{
		Id:          0,
		Title:       r.FormValue("title"),
		TitleOrigin: r.FormValue("title-origin"),
		Author:      FilterStringBySep(r.FormValue("author"), ","),
		Translator:  FilterStringBySep(r.FormValue("translator"), ","),
		Publisher:   r.FormValue("publisher"),
		Cover:       r.FormValue("cover"),
		Tag:         FilterStringBySep(r.FormValue("tag"), ","),
		Reads:       nil,
		Likes:       nil,
	}
	err := db.AddBook(uuid, book)
	if err != nil {
		Fprint(w, ResponseER)
		return
	}
	Fprint(w, ResponseOK)
}

func bookLike(w http.ResponseWriter, r *http.Request) {
	uuid := session.GetUUID(r)
	if uuid <= 0 {
		return
	}
	bid := Int64(r.FormValue("b"))
	if bid <= 0 {
		return
	}
	likes := db.GetBookLikes(uuid, bid)
	Fprint(w, Marshal(likes))
}

func bookReadStart(w http.ResponseWriter, r *http.Request) {
	uuid := session.GetUUID(r)
	if uuid <= 0 {
		return
	}
	bid := Int64(r.FormValue("bid"))
	if bid <= 0 {
		return
	}
	if db.AddBookReadStart(uuid, bid) != nil {
		Fprint(w, ResponseER)
		return
	}
	Fprint(w, ResponseOK)
}

func bookReadFinish(w http.ResponseWriter, r *http.Request) {
	uuid := session.GetUUID(r)
	if uuid <= 0 {
		return
	}
	bid := Int64(r.FormValue("bid"))
	if bid <= 0 {
		return
	}
	if db.AddBookReadFinish(uuid, bid) != nil {
		Fprint(w, ResponseER)
		return
	}
	Fprint(w, ResponseOK)
}

func bookLikeAdd(w http.ResponseWriter, r *http.Request) {
	uuid := session.GetUUID(r)
	if uuid <= 0 {
		return
	}
	bid := Int64(r.FormValue("bid"))
	if bid <= 0 {
		return
	}
	page := Int64(r.FormValue("page"))
	content := strings.TrimSpace(r.FormValue("content"))
	comment := strings.TrimSpace(r.FormValue("comment"))
	if db.AddBookLike(uuid, bid, page, content, comment) != nil {
		Fprint(w, ResponseER)
		return
	}
	Fprint(w, ResponseOK)
}

func bookLikeDel(w http.ResponseWriter, r *http.Request) {
	uuid := session.GetUUID(r)
	if uuid <= 0 {
		return
	}
	id := Int64(r.FormValue("id"))
	if id <= 0 {
		return
	}
	if db.DeleteBookLike(uuid, id) != nil {
		Fprint(w, ResponseER)
		return
	}
	Fprint(w, ResponseOK)
}

func bookLikeFix(w http.ResponseWriter, r *http.Request) {
	uuid := session.GetUUID(r)
	if uuid <= 0 {
		return
	}
	id := Int64(r.FormValue("id"))
	if id <= 0 {
		return
	}
	page := Int64(r.FormValue("page"))
	content := strings.TrimSpace(r.FormValue("content"))
	comment := strings.TrimSpace(r.FormValue("comment"))
	if db.UpdateBookLike(uuid, id, page, content, comment) != nil {
		Fprint(w, ResponseER)
		return
	}
	Fprint(w, ResponseOK)
}
