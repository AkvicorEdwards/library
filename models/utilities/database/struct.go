package database

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
)

type TableCategories struct {
	Id            int64  `json:"id"`
	CategoryName  string `json:"category_name"`
	CoverImgSrc   string `json:"cover_img_src"`
	NumInCategory int64  `json:"num_in_category"`
	HitCount      int64  `json:"hit_count"`
}

type TableCategoriesSQL struct {
	Id            sql.NullInt64  `json:"id"`
	CategoryName  sql.NullString `json:"category_name"`
	CoverImgSrc   sql.NullString `json:"cover_img_src"`
	NumInCategory sql.NullInt64  `json:"num_in_category"`
	HitCount      sql.NullInt64  `json:"hit_count"`
}

func (a *TableCategoriesSQL) Transfer() (b TableCategories) {
	b.Id = a.Id.Int64
	b.CategoryName = base64D2S(a.CategoryName.String)
	b.CoverImgSrc = base64D2S(a.CoverImgSrc.String)
	b.NumInCategory = a.NumInCategory.Int64
	b.HitCount = a.HitCount.Int64
	return
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

type TableUser struct {
	Id         int64
	Username   string
	Password   string
	NickName   string
	Email      string
	Permission int64
}

type TableUserSQL struct {
	Id         sql.NullInt64
	Username   sql.NullString
	Password   sql.NullString
	NickName   sql.NullString
	Email      sql.NullString
	Permission sql.NullInt64
}

func (a *TableUserSQL) Transfer() (b TableUser) {
	b.Id = a.Id.Int64
	b.Username = a.Username.String
	b.Password = a.Password.String
	b.NickName = a.NickName.String
	b.Email = a.Email.String
	b.Permission = a.Permission.Int64
	return
}

func base64D2S(str string) string {
	b, _ := base64.StdEncoding.DecodeString(str)
	return string(b)
}

func base64D2B(str string) []byte {
	b, _ := base64.StdEncoding.DecodeString(str)
	return b
}
