package book

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"library/config"
	"library/mysql"
	"library/tool"
	"library/tpl"
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
		maxValueBytes := int64(2<<20)
		form := make(map[string]interface{}, 0)
		form["book"] = ""
		form["author"] = ""
		form["translator"] = ""
		form["publisher"] = ""
		form["cover"] = ""
		form["tag"] = make([]string,0)
		form["time"] = make([]mysql.FormTime, 0)
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
				switch name {
				case "tag":
					form["tag"] = tools.StringSplitBySpace(b.String())
				case "time":
					form["time"] = []mysql.FormTime{
						{
							Start: b.String(),
							End:   "2199-12-05",
						},
					}
				default:
					if _, ok := form[name]; ok {
						form[name] = b.String()
					} else {
						fmt.Println("Unhandled form name:",name)
					}
				}
				continue
			}
			fileName := strconv.FormatInt(time.Now().Unix(), 10)+ "_" + strconv.FormatInt(int64(rand.Int()), 10) + "_" + part.FileName()
			form["cover"] = fileName
			func(){
				dst, err := os.Create(config.Path.Cover + fileName)
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
		tag, err := json.Marshal(form["tag"])
		tm, err := json.Marshal(form["time"])
		favour, err := json.Marshal(make([]string, 0))
		_, err = mysql.Exec("INSERT INTO books (book, author, translator, publisher, cover, tag, reading, favour) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			form["book"], form["author"], form["translator"], form["publisher"], form["cover"], string(tag), string(tm), string(favour))

		if err != nil {
			_, _ = mysql.Exec("alter table books AUTO_INCREMENT=1")
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, r.Referer(), http.StatusFound)
		return
	}
	phrase := map[string]interface{}{
		"lang": "zh-cn",
		"title": "Add",
		"Time": time.Now().Format("2006-01-02"),
	}

	if err := tpl.BookAdd.Execute(w, phrase); err != nil {
		fmt.Println(err)
		_, _ = fmt.Fprintf(w, "%v", "Error")
	}
}
func AddFavour(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.PostFormValue("id")
		page := r.PostFormValue("page")
		tm := time.Now().Format("2006-01-02 15:04:05")
		content := r.PostFormValue("contents")
		comment := r.PostFormValue("comment")
		_, err := mysql.Exec(`UPDATE books SET favour = JSON_ARRAY_APPEND(favour, '$', JSON_OBJECT("page", ?, "time", ?, "content", ?, "comment", ?)) WHERE id = ?`, page, tm, content, comment, id)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = mysql.Exec(`UPDATE books SET favour_cnt=favour_cnt+1 WHERE id = ?`, id)
		if err != nil {
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}
func SetTime(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ids := r.PostFormValue("ids")
		id, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			fmt.Println(err)
			return
		}
		typ := `'$[`+fmt.Sprint(id-1)+`].`+r.PostFormValue("typ")+`_time'`
		tm := r.PostFormValue("time")
		_, err = mysql.Exec(`UPDATE books SET reading = JSON_SET(reading, `+typ+`, ?) WHERE id = ?`, tm, ids)

		if err != nil {
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}
func SetStartRead(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ids := r.PostFormValue("ids")
		tm := r.PostFormValue("time")
		_, err := mysql.Exec(`UPDATE books SET reading = JSON_ARRAY_APPEND(reading, '$', JSON_OBJECT("start_time", ?, "end_time", "2199-12-05")) WHERE id = ?`, tm, ids)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = mysql.Exec(`UPDATE books SET reading_cnt=reading_cnt+1 WHERE id = ?`, ids)
		if err != nil {
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}
func FixCover(w http.ResponseWriter, r *http.Request) {
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
					_, _ = fmt.Fprintln(w, err)
				}
				ids = b.String()
				continue
			}
			if len(name) == 0 || len(part.FileName()) == 0{
				continue
			}

			cover = strconv.FormatInt(time.Now().Unix(), 10)+ "_" + strconv.FormatInt(int64(rand.Int()), 10) + "_" + part.FileName()
			func(){
				dst, err := os.Create(config.Path.Cover + cover)
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
			_, _ = fmt.Fprint(w, "No file")
			return
		}
		_, err = mysql.Exec(`UPDATE books SET cover = ? WHERE id = ?`, cover, ids)

		if err != nil {
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}
func Fix(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ids := r.PostFormValue("ids")
		book := r.PostFormValue("book")
		author := r.PostFormValue("author")
		translator := r.PostFormValue("translator")
		publisher := r.PostFormValue("publisher")
		tag, err := json.Marshal(tools.StringSplitBySpace(r.PostFormValue("tag")))
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = mysql.Exec(`UPDATE books SET book=?, author=?, translator=?, publisher=?, tag=? WHERE id = ?`, book, author, translator, publisher, string(tag), ids)
		if err != nil {
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}
func Index(w http.ResponseWriter, r *http.Request) {
	sqlStr := fmt.Sprintf("SELECT id, book, author, translator, tag, reading, reading_cnt, favour_cnt FROM books")

	data := func() (data []mysql.TableBooks) {
		rows, err := mysql.Query(sqlStr)
		if err != nil {
			panic(err)
		}
		row := mysql.TableBooksSQL{}
		for rows.Next() {
			if err = rows.Scan(&row.Id, &row.Book, &row.Author, &row.Translator, &row.Tag, &row.Reading, &row.ReadingCnt, &row.FavourCnt);
				err != nil {
				panic(err)
				return
			}
			data = append(data, row.Transfer())
		}
		return
	}()
	book, _ := json.Marshal(data)
	urlShort := fmt.Sprintf("%s://%s/", config.Server.Protocol, r.Host)
	phrase := map[string]interface{}{
		"lang": "zh-cn",
		"title": "Books",
		"Books": string(book),
		"UrlShort": urlShort,
	}

	if err := tpl.BookIndex.Execute(w, phrase); err != nil {
		fmt.Println(err)
		_, _ = fmt.Fprintf(w, "%v", "Error")
	}
}
func View(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(fmt.Sprintf("%s", r.URL)[6:], "?")[0]
	sqlStr := fmt.Sprintf("SELECT * FROM books WHERE id = ? LIMIT 1")

	data := func() (data mysql.TableBooks) {
		rows, err := mysql.Query(sqlStr, id)
		if err != nil {
			panic(err)
		}
		row := mysql.TableBooksSQL{}
		if rows.Next() {
			if err = rows.Scan(&row.Id, &row.Book, &row.Author, &row.Translator, &row.Publisher, &row.Cover, &row.Tag, &row.Reading, &row.ReadingCnt, &row.Favour, &row.FavourCnt);
				err != nil {
				panic(err)
				return
			}
			return row.Transfer()
		}
		return
	}()

	book, _ := json.Marshal(data)
	urlShort := fmt.Sprintf("%s://%s/", config.Server.Protocol, r.Host)
	phrase := map[string]interface{}{
		"lang": "zh-cn",
		"title": "Books",
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
