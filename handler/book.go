package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"library/def"
	"library/operator"
	"library/tpl"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		mr, err := r.MultipartReader()
		if err != nil {
			fmt.Println(err)
			return
		}
		maxValueBytes := int64(2 << 20)
		form := make(map[string]string, 0)
		form["book"] = ""
		form["author"] = ""
		form["translator"] = ""
		form["publisher"] = ""
		form["cover"] = ""
		form["tag"] = ""
		form["time"] = ""

		for {
			part, err := mr.NextPart()
			if err != nil {
				break
			}
			name := part.FormName()
			if len(name) == 0 {
				continue
			}
			if len(part.FileName()) == 0 || name != "cover" {
				var b bytes.Buffer
				_, err := io.CopyN(&b, part, maxValueBytes)
				if err != nil && err != io.EOF {
					fmt.Println(err)
					_, _ = fmt.Fprintln(w, err)
				}
				if _, ok := form[name]; ok {
					form[name] = b.String()
				} else {
					fmt.Println("Unhandled form name:", name)
				}
				continue
			}
			fileName := strconv.FormatInt(time.Now().Unix(), 10) + "_" + strconv.FormatInt(int64(rand.Int()), 10) + "_" + part.FileName()
			form["cover"] = fileName
			func() {
				dst, err := os.Create("./cover/" + fileName)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer dst.Close()
				for {
					buffer := make([]byte, 100000)
					cBytes, errs := part.Read(buffer)

					n, err := dst.Write(buffer[0:cBytes])
					if err != nil {
						fmt.Println("File copy Field")
						fmt.Println(cBytes)
						fmt.Println(n)
						fmt.Println(err)
						break
					}
					if errs == io.EOF {
						break
					}
				}
			}()
		}
		res := operator.AddBook(form["book"], form["author"], form["translator"], form["publisher"], form["cover"], form["tag"], form["time"])

		if !res {
			log.Println(err)
			return
		}
		http.Redirect(w, r, r.Referer(), http.StatusFound)
		return
	}

	phrase := map[string]interface{}{
		"Time":  time.Now().Format("2006-01-02"),
	}

	if err := tpl.BookAdd.Execute(w, phrase); err != nil {
		log.Println(err)
		Fprintf(w, "%v", "Error")
	}
}

func bookIndex(w http.ResponseWriter, r *http.Request) {
	book, err := json.Marshal(operator.GetBookList())
	if err != nil {
		log.Println(err)
		return
	}
	urlShort := fmt.Sprintf("%s://%s/", def.Protocol, r.Host)
	phrase := map[string]interface{}{
		"lang": "zh-cn",
		"title": "Books",
		"Books": string(book),
		"UrlShort": urlShort,
	}

	if err := tpl.BookIndex.Execute(w, phrase); err != nil {
		log.Println(err)
		_, _ = fmt.Fprintf(w, "%v", "Error")
	}
}

func bookView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.Split(fmt.Sprintf("%s", r.URL)[6:], "?")[0])
	if err != nil {
		log.Println(err)
		return
	}
	data := operator.GetBook(id)
	book, _ := json.Marshal(data)
	urlShort := fmt.Sprintf("%s://%s/", def.Protocol, r.Host)
	if id == 1{
		phrase := map[string]interface{}{
			"lang": "zh-cn",
			"title": data.Book,
			"Books": string(book),
			"UrlShort": urlShort,
			"Id": id,
		}
		if err := tpl.BookComplex.Execute(w, phrase); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
		return
	}
	phrase := map[string]interface{}{
		"lang": "zh-cn",
		"title": data.Book,
		"Books": string(book),
		"UrlShort": urlShort,
		"Id": id,
		"ReadCnt": data.ReadingCnt,
		"Time": time.Now().Format("2006-01-02"),

		"FixBook": data.Book,
		"FixAuthor": data.Author,
		"FixTranslator": data.Translator,
		"FixPublisher": data.Publisher,
		"FixTag": func() (str string) {
			for _, v := range data.Tag {
				str += v + " "
			}
			return
		}(),
	}
	if err := tpl.Book.Execute(w, phrase); err != nil {
		fmt.Println(err)
		_, _ = fmt.Fprintf(w, "%v", "Error")
	}
}

func bookAddFavour(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			return
		}
		page := r.PostFormValue("page")
		tm := time.Now().Format("2006-01-02 15:04:05")
		content := r.PostFormValue("contents")
		comment := r.PostFormValue("comment")
		operator.AddFavour(id, page, tm, content, comment)
		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}

func bookSetTime(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.PostFormValue("ids"))
		if err != nil {
			Fprint(w,"Error")
			return
		}
		idInJson, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			Fprint(w,"Error")
			return
		}
		ok := operator.SetTime(id, idInJson, r.PostFormValue("typ"),
			r.PostFormValue("time"))
		if !ok {
			Fprint(w,"Error")
		}
		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}

func bookSetStartRead(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.PostFormValue("ids"))
		if err != nil {
			Fprint(w,"Error")
			return
		}
		tm := r.PostFormValue("time")
		ok := operator.SetStartRead(id, tm)
		if !ok {
			Fprint(w,"Error")
			return
		}
		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}

func bookFix(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PostFormValue("ids"))
	if err != nil {
		Fprint(w, "Error")
	}
	book := r.PostFormValue("book")
	author := r.PostFormValue("author")
	translator := r.PostFormValue("translator")
	publisher := r.PostFormValue("publisher")
	tag := r.PostFormValue("tag")
	ok := operator.Fix(id, book, author, translator, publisher, tag)
	if !ok {
		Fprint(w,"Error")
		return
	}
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}

func bookFixCover(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		mr, err := r.MultipartReader()
		if err != nil {
			fmt.Println(err)
			return
		}
		cover := ""
		ids := ""
		for {
			part, err := mr.NextPart()
			if err != nil {
				break
			}
			name := part.FormName()
			if name == "ids" {
				var b bytes.Buffer
				_, err := io.CopyN(&b, part, 1000000)
				if err != nil && err != io.EOF {
					fmt.Println(err)
					Fprintln(w, err)
				}
				ids = b.String()
				continue
			}
			if len(name) == 0 || len(part.FileName()) == 0{
				continue
			}

			cover = strconv.FormatInt(time.Now().Unix(), 10)+ "_" + strconv.FormatInt(int64(rand.Int()), 10) + "_" + part.FileName()
			func(){
				dst, err := os.Create("./cover/" + cover)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer dst.Close()
				for {
					buffer := make([]byte, 100000)
					cBytes, errs := part.Read(buffer)

					n, err := dst.Write(buffer[0:cBytes])
					if err != nil {
						fmt.Println("File copy Field")
						fmt.Println(cBytes)
						fmt.Println(n)
						fmt.Println(err)
						break
					}
					if errs == io.EOF {
						break
					}
				}
			}()
		}
		if len(cover) == 0 {
			Fprint(w, "No file")
			return
		}

		id, err := strconv.Atoi(ids)
		if err != nil {
			Fprint(w, "Error")
			return
		}

		if !operator.FixCover(id, cover) {
			Fprint(w, "Error")
			return
		}

		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}

func bookFixFavour(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PostFormValue("id"))
	if err != nil {
		Fprint(w, "Error")
		return
	}
	ids, err := strconv.Atoi(r.PostFormValue("ids"))
	if err != nil {
		Fprint(w, "Error")
		return
	}
	page := r.PostFormValue("page")
	tm := r.PostFormValue("time")
	content := r.PostFormValue("contents")
	comment := r.PostFormValue("comment")

	ok := operator.FixFavour(id, ids, page, tm, content, comment)
	if !ok {
		Fprint(w, "Error")
		return
	}
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}

func bookDelFavour(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			Fprint(w, "Error")
			return
		}
		ids, err := strconv.Atoi(r.PostFormValue("ids"))
		if err != nil {
			Fprint(w, "Error")
			return
		}
		if !operator.DelFavour(id, ids) {
			Fprint(w, "Error")
			return
		}
		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}
