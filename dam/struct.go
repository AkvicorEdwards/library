package dam

import "encoding/json"

type Inc struct {
	Name string // 表名
	Val  int    // 自增值
}

type User struct {
	Uuid     int    `json:"uuid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type DurTime struct {
	Start string `json:"start_time"`
	End   string `json:"end_time"`
}

type Favour struct {
	Page    string `json:"page"`
	Time    string `json:"time"`
	Content string `json:"content"`
	Comment string `json:"comment"`
}

type Book struct {
	Id         int       `json:"id"`
	Book       string    `json:"book"`
	Author     string    `json:"author"`
	Translator string    `json:"translator"`
	Publisher  string    `json:"publisher"`
	Cover      string    `json:"cover"`
	Tag        []string  `json:"tag"`
	Reading    []DurTime `json:"reading"`
	ReadingCnt int64     `json:"reading_cnt"`
	Favour     []Favour  `json:"favour"`
	FavourCnt  int64     `json:"favour_cnt"`
}

func (b *Book) Transfer() (book BookDB) {
	tag, err := json.Marshal(b.Tag)
	if err != nil {
		return
	}
	reading, err := json.Marshal(b.Reading)
	if err != nil {
		return
	}
	favour, err := json.Marshal(b.Favour)
	if err != nil {
		return
	}

	book = BookDB{
		Id:         b.Id,
		Book:       b.Book,
		Author:     b.Author,
		Translator: b.Translator,
		Publisher:  b.Publisher,
		Cover:      b.Cover,
		Tag:        string(tag),
		Reading:    string(reading),
		ReadingCnt: b.FavourCnt,
		Favour:     string(favour),
		FavourCnt:  b.FavourCnt,
	}
	return
}

type BookDB struct {
	Id         int    `json:"id"`
	Book       string `json:"book"`
	Author     string `json:"author"`
	Translator string `json:"translator"`
	Publisher  string `json:"publisher"`
	Cover      string `json:"cover"`
	Tag        string `json:"tag"`
	Reading    string `json:"reading"`
	ReadingCnt int64  `json:"reading_cnt"`
	Favour     string `json:"favour"`
	FavourCnt  int64  `json:"favour_cnt"`
}

func (b *BookDB) Transfer() (book Book) {
	book = Book{
		Id:         b.Id,
		Book:       b.Book,
		Author:     b.Author,
		Translator: b.Translator,
		Publisher:  b.Publisher,
		Cover:      b.Cover,
		Tag:        make([]string, 0),
		Reading:    make([]DurTime, 0),
		ReadingCnt: b.ReadingCnt,
		Favour:     make([]Favour, 0),
		FavourCnt:  b.FavourCnt,
	}
	err := json.Unmarshal([]byte(b.Tag), &book.Tag)
	if err != nil {
		return Book{}
	}
	err = json.Unmarshal([]byte(b.Reading), &book.Reading)
	if err != nil {
		return Book{}
	}
	err = json.Unmarshal([]byte(b.Favour), &book.Favour)
	if err != nil {
		return Book{}
	}
	return
}

type BookDBBriefing struct {
	Id         int    `json:"id"`
	Book       string `json:"book"`
	Author     string `json:"author"`
	Translator string `json:"translator"`
	Publisher  string `json:"publisher"`
	Cover      string `json:"cover"`
	Tag        string `json:"tag"`
	ReadingCnt int64  `json:"reading_cnt"`
	FavourCnt  int64  `json:"favour_cnt"`
}

func (b *BookDBBriefing) Transfer() (book Book) {
	book = Book{
		Id:         b.Id,
		Book:       b.Book,
		Author:     b.Author,
		Translator: b.Translator,
		Publisher:  b.Publisher,
		Cover:      b.Cover,
		Tag:        make([]string, 0),
		Reading:    make([]DurTime, 0),
		ReadingCnt: b.ReadingCnt,
		Favour:     make([]Favour, 0),
		FavourCnt:  b.FavourCnt,
	}
	err := json.Unmarshal([]byte(b.Tag), &book.Tag)
	if err != nil {
		return Book{}
	}
	return
}