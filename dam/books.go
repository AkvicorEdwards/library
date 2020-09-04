package dam

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"log"
)

func AddBook(book Book) bool {
	if !Connected {
		Connect()
	}
	lockBooks.Lock()
	defer lockBooks.Unlock()

	book.Id = GetInc("books") + 1
	bookT := book.Transfer()

	if err := db.Table("books").Create(&bookT).Error; err != nil {
		log.Println(err)
		return false
	}

	UpdateInc("books", book.Id)

	return true
}

func GetBookList() []Book {
	if !Connected {
		Connect()
	}
	booksDBBriefing := make([]BookDBBriefing, 0)
	books := make([]Book, 0)
	if err := db.Table("books").Find(&booksDBBriefing).Error; err != nil {
		log.Println(err)
		return books
	}
	for _, v := range booksDBBriefing {
		books = append(books, v.Transfer())
	}
	return books
}

func GetBook(id int) Book {
	if !Connected {
		Connect()
	}
	book := BookDB{}
	if err := db.Table("books").Where("id=?", id).First(&book).Error; err != nil {
		return Book{}
	}
	return book.Transfer()
}

func UpdateFavour(id int, favour []Favour) bool {
	if !Connected {
		Connect()
	}
	fav, err := json.Marshal(favour)
	if err != nil {
		return false
	}
	err = db.Table("books").Where("id=?", id).Updates(map[string]interface{}{
		"favour": string(fav),
	}).Error

	if err != nil {
		return false
	}

	return true
}

func UpdateTime(id int, durTime []DurTime) bool {
	if !Connected {
		Connect()
	}

	dur, err := json.Marshal(durTime)
	if err != nil {
		return false
	}
	err = db.Table("books").Where("id=?", id).Updates(map[string]interface{}{
		"reading": string(dur),
	}).Error

	if err != nil {
		return false
	}

	return true
}

func UpdateBookInfo(book Book) bool {
	if !Connected {
		Connect()
	}
	bookT := book.Transfer()
	err := db.Table("books").Where("id=?", bookT.Id).Updates(map[string]interface{}{
		"book": bookT.Book,
		"author": bookT.Author,
		"translator": bookT.Translator,
		"publisher": bookT.Publisher,
		"tag": bookT.Tag,
	}).Error

	if err != nil {
		return false
	}

	return true
}

func UpdateBookCover(id int, cover string) bool {
	if !Connected {
		Connect()
	}

	err := db.Table("books").Where("id=?", id).Updates(map[string]interface{}{
		"cover": cover,
	}).Error

	if err != nil {
		return false
	}

	return true
}

func UpReadingCnt(id int) bool {
	if !Connected {
		Connect()
	}

	err := db.Table("books").Where("id=?", id).Update("reading_cnt",
		gorm.Expr("reading_cnt+1")).Error
	if err != nil {
		return false
	}

	return true
}
