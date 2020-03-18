package mysql

import (
	"database/sql"
	"encoding/json"
)

type TableUser struct {
	Id			int64
	UserName	string
	UserPwd		string
	NickName	string
	Email		string
	Permissions	int64
}

type TableUserSQL struct {
	Id			sql.NullInt64
	UserName	sql.NullString
	UserPwd		sql.NullString
	NickName	sql.NullString
	Email		sql.NullString
	Permissions	sql.NullInt64
}

func (a *TableUserSQL) Transfer() (b TableUser) {
	b.Id 			= a.Id.Int64
	b.UserName 		= a.UserName.String
	b.UserPwd 		= a.UserPwd.String
	b.NickName 		= a.NickName.String
	b.Email 		= a.Email.String
	b.Permissions	= a.Permissions.Int64
	return
}

func (TableUserSQL) TableName() string {
	return "users"
}
func (TableUser) TableName() string {
	return "users"
}

type FormTime struct {
	Start string `json:"start_time"`
	End string `json:"end_time"`
}
type FormFavour struct {
	Page string `json:"page"`
	Time string `json:"time"`
	Content string `json:"content"`
	Comment string `json:"comment"`
}

type TableBooks struct {
	Id         int64    `json:"id"`
	Book       string   `json:"book"`
	Author     string   `json:"author"`
	Translator string   `json:"translator"`
	Publisher  string   `json:"publisher"`
	Cover      string   `json:"cover"`
	Tag        []string `json:"tag"`
	Reading    []FormTime `json:"reading"`
	ReadingCnt int64    `json:"reading_cnt"`
	Favour     []FormFavour `json:"favour"`
	FavourCnt  int64    `json:"favour_cnt"`
}

type TableBooksSQL struct {
	Id         sql.NullInt64  `json:"id"`
	Book       sql.NullString `json:"book"`
	Author     sql.NullString `json:"author"`
	Translator sql.NullString `json:"translator"`
	Publisher  sql.NullString `json:"publisher"`
	Cover      sql.NullString `json:"cover"`
	Tag        sql.NullString `json:"tag"`
	Reading    sql.NullString `json:"reading"`
	ReadingCnt sql.NullInt64  `json:"reading_cnt"`
	Favour     sql.NullString `json:"favour"`
	FavourCnt  sql.NullInt64  `json:"favour_cnt"`
}

func (a *TableBooksSQL) Transfer() (b TableBooks) {
	b.Id = a.Id.Int64
	b.Book = a.Book.String
	b.Author = a.Author.String
	b.Translator = a.Translator.String
	b.Publisher = a.Publisher.String
	b.Cover = a.Cover.String
	_ = json.Unmarshal([]byte(a.Tag.String), &b.Tag)
	_ = json.Unmarshal([]byte(a.Reading.String), &b.Reading)
	b.ReadingCnt = a.ReadingCnt.Int64
	_ = json.Unmarshal([]byte(a.Favour.String), &b.Favour)
	b.FavourCnt = a.FavourCnt.Int64
	return
}

func (TableBooks) TableName() string {
	return "books"
}
func (TableBooksSQL) TableName() string {
	return "books"
}