package db

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Bookshelf 表，以Book为项
type Bookshelf struct {
	UUID         int64  `json:"uuid"`
	Username     string `json:"username"`
	ProfilePhoto string `json:"profile_photo"`
	Nickname     string `json:"nickname"`
	Books        []Book   `json:"books"`
	Likes        []Like `json:"likes"`
}

// Bookshelf 表的项
type Book struct {
	Id          int64      `json:"id"`
	Title       string     `json:"title"`
	TitleOrigin string     `json:"title_origin"`
	Author      []string   `json:"author"`
	Translator  []string   `json:"translator"`
	Publisher   string     `json:"publisher"`
	Cover       string     `json:"cover"`
	Tag         []string   `json:"tag"`
	Reads       []BookRead `json:"reads"`
	Likes       []BookLike `json:"likes"`
}

func (b *Book) AuthorMarshal() string {
	if b.Author == nil {
		return `[]`
	}
	str, err := json.Marshal(b.Author)
	if err != nil {
		return `[]`
	}
	return string(str)
}

func (b *Book) TranslatorMarshal() string {
	if b.Translator == nil {
		return `[]`
	}
	str, err := json.Marshal(b.Translator)
	if err != nil {
		return `[]`
	}
	return string(str)
}

func (b *Book) TagMarshal() string {
	if b.Tag == nil {
		return `[]`
	}
	str, err := json.Marshal(b.Tag)
	if err != nil {
		return `[]`
	}
	return string(str)
}

func (b *Book) ReadsMarshal() string {
	if b.Reads == nil {
		return `[]`
	}
	str, err := json.Marshal(b.Reads)
	if err != nil {
		return `[]`
	}
	return string(str)
}

func (b *Book) LikesMarshal() string {
	if b.Likes == nil {
		return `[]`
	}
	str, err := json.Marshal(b.Likes)
	if err != nil {
		return `[]`
	}
	return string(str)
}

type BookRead struct {
	Id    int64   `json:"id"`
	Bid    int64   `json:"bid"`
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type BookLike struct {
	Id      int64    `json:"id"`
	Bid    int64   `json:"bid"`
	Page    int64  `json:"page"`
	Time    int64  `json:"time"`
	Content string `json:"content"`
	Comment string `json:"comment"`
}

type Like struct {
	Id      int64  `json:"id"`
	Origin  string `json:"origin"`
	Time    int64  `json:"time"`
	Content string `json:"content"`
	Comment string `json:"comment"`
}

func AddBook(uuid int64, book Book) error {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("book_%d", uuid)
	book.Id = GetInc(table) + 1
	err := db.Table(table).Create(map[string]interface{}{
		"id": book.Id,
		"title": book.Title,
		"title_origin": book.TitleOrigin,
		"author": book.AuthorMarshal(),
		"translator": book.TranslatorMarshal(),
		"publisher": book.Publisher,
		"cover": book.Cover,
		"tag": book.TagMarshal(),
	}).Error
	if err != nil {
		return err
	}
	updateInc(table, book.Id)
	return nil
}

func UpdateBook(uuid int64, book Book) error {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("book_%d", uuid)
	err := db.Table(table).Where("id=?", book.Id).Updates(map[string]interface{}{
		"title": book.Title,
		"title_origin": book.TitleOrigin,
		"author": book.AuthorMarshal(),
		"translator": book.TranslatorMarshal(),
		"publisher": book.Publisher,
		"cover": book.Cover,
		"tag": book.TagMarshal(),
	}).Error
	return err
}

func AddBookLike(uuid, bid, page int64, content, comment string) error {
	if !Connected {
		Connect()
	}
	if !BidValid(uuid, bid) {
		return ErrorBidNotValid
	}
	table := fmt.Sprintf("book_like_%d", uuid)
	err := db.Table(table).Create(map[string]interface{}{
		"bid": bid,
		"page": page,
		"time": time.Now().Unix(),
		"content": content,
		"comment": comment,
	}).Error
	return err
}

func UpdateBookLike(uuid, id, page int64, content, comment string) error {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("book_like_%d", uuid)
	err := db.Table(table).Where("id=?", id).UpdateColumns(map[string]interface{}{
		"page": page,
		"content": content,
		"comment": comment,
	}).Error
	return err
}

func AddBookReadStart(uuid, bid int64) error {
	if !Connected {
		Connect()
	}
	if !BidValid(uuid, bid) {
		return ErrorBidNotValid
	}
	last := GetBookLastRead(uuid, bid)
	if last != nil && last.End <= 0 {
		return ErrorUnfinished
	}
	table := fmt.Sprintf("book_read_%d", uuid)
	err := db.Table(table).Create(map[string]interface{}{
		"bid": bid,
		"start": time.Now().Unix(),
		"end": 0,
	}).Error
	return err
}

func AddBookReadFinish(uuid, bid int64) error {
	if !Connected {
		Connect()
	}
	if !BidValid(uuid, bid) {
		return ErrorBidNotValid
	}
	last := GetBookLastRead(uuid, bid)
	if last == nil {
		return ErrorNoBook
	}
	if last.End > 0 {
		return ErrorFinished
	}
	table := fmt.Sprintf("book_read_%d", uuid)
	err := db.Table(table).Where("id=?", last.Id).Update("end", time.Now().Unix()).Error

	return err
}

func GetBookLastRead(uuid, bid int64) *BookRead {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("book_read_%d", uuid)
	read := &BookRead{}
	err := db.Table(table).Where("bid=? AND del=0", bid).Last(read).Error
	if err != nil {
		return nil
	}
	return read
}

type book struct {
	Id int64
	Title string
	TitleOrigin string
	Author string
	Translator string
	Publisher string
	Cover string
	Tag string
}

func (b *book) UnMarshal() Book {
	res := Book{
		Id:          b.Id,
		Title:       b.Title,
		TitleOrigin: b.TitleOrigin,
		Author:      make([]string, 0),
		Translator:  make([]string, 0),
		Publisher:   b.Publisher,
		Cover:       b.Cover,
		Tag:         make([]string, 0),
		Reads:       nil,
		Likes:       nil,
	}
	err := json.Unmarshal([]byte(b.Author), &res.Author)
	if err != nil { return Book{} }
	err = json.Unmarshal([]byte(b.Translator), &res.Translator)
	if err != nil { return Book{} }
	err = json.Unmarshal([]byte(b.Tag), &res.Tag)
	if err != nil { return Book{} }
	return res
}

func GetBookLikes(uuid, bid int64) []BookLike {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("book_like_%d", uuid)
	res := make([]BookLike, 0)
	err := db.Table(table).Where("bid=? AND del=0", bid).Find(&res).Error
	if err != nil {
		return make([]BookLike, 0)
	}
	return res
}

func GetBookReads(uuid, bid int64) []BookRead {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("book_read_%d", uuid)
	res := make([]BookRead, 0)
	err := db.Table(table).Where("bid=? AND del=0", bid).Find(&res).Error
	if err != nil {
		return make([]BookRead, 0)
	}
	return res
}

func GetBooks(uuid int64) []Book {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("book_%d", uuid)
	res := make([]book, 0)
	err := db.Table(table).Find(&res).Error
	if err != nil {
		return nil
	}

	books := make([]Book, 0, len(res))
	for k, v := range res {
		books = append(books, v.UnMarshal())
		books[k].Likes = GetBookLikes(uuid, v.Id)
		books[k].Reads = GetBookReads(uuid, v.Id)
	}
	return books
}

func GetBookList(uuid int64) []Book {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("book_%d", uuid)
	res := make([]book, 0)
	err := db.Table(table).Find(&res).Error
	if err != nil {
		return nil
	}

	books := make([]Book, 0, len(res))
	for _, v := range res {
		books = append(books, v.UnMarshal())
	}
	return books
}

func GetBookInfo(uuid, bid int64) Book {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("book_%d", uuid)
	res := book{}
	err := db.Table(table).Where("id=?", bid).First(&res).Error
	if err != nil {
		log.Println(err)
		return Book{}
	}
	re := res.UnMarshal()
	return re
}

func GetBook(uuid, bid int64) Book {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("book_%d", uuid)
	res := book{}
	err := db.Table(table).Where("id=?", bid).First(&res).Error
	if err != nil {
		log.Println(err)
		return Book{}
	}
	re := res.UnMarshal()
	re.Likes = GetBookLikes(uuid, res.Id)
	re.Reads = GetBookReads(uuid, res.Id)
	return re
}

func GetLikes(uuid int64) []Like {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("like_%d", uuid)
	likes := make([]Like, 0)
	err := db.Table(table).Where("del=0").Find(&likes).Error
	if err != nil {
		return nil
	}
	return likes
}

func GetBookShelf(uuid int64) Bookshelf {
	if !Connected {
		Connect()
	}
	user := GetUserInfo(uuid)
	bookshelf := Bookshelf{
		UUID:         user.UUID,
		Username:     user.Username,
		ProfilePhoto: user.ProfilePhoto,
		Nickname:     user.Nickname,
		Books:        GetBooks(uuid),
		Likes:        GetLikes(uuid),
	}
	return bookshelf
}

func AddLike(uuid int64, origin, content, comment string) error {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("like_%d", uuid)
	err := db.Table(table).Create(map[string]interface{}{
		"origin": origin,
		"time": time.Now().Unix(),
		"content": content,
		"comment": comment,
	}).Error
	return err
}

func DeleteLike(uuid, id int64) error {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("like_%d", uuid)
	err := db.Table(table).Where("id=?", id).Update("del", time.Now().Unix()).Error
	return err
}

func UpdateLike(uuid, id int64, origin, content, comment string) error {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("like_%d", uuid)
	err := db.Table(table).Where("id=?", id).Updates(map[string]interface{}{
		"origin": origin,
		"content": content,
		"comment": comment,
	}).Error
	return err
}

func DeleteBookLike(uuid, id int64) error {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("book_like_%d", uuid)
	err := db.Table(table).Where("id=?", id).Update("del", time.Now().Unix()).Error
	return err
}

func DeleteBookRead(uuid, id int64) error {
	if !Connected {
		Connect()
	}
	table := fmt.Sprintf("book_read_%d", uuid)
	err := db.Table(table).Where("id=?", id).Update("del", time.Now().Unix()).Error
	return err
}


