package operator

import (
	"library/dam"
	"library/util"
)

func AddBook(book, author, translator, publisher, cover, tag, readingStart string) bool {
	bookS := dam.Book{
		Id:         0,
		Book:       book,
		Author:     author,
		Translator: translator,
		Publisher:  publisher,
		Cover:      cover,
		Tag:        util.StringSplitBySpace(tag),
		Reading:    make([]dam.DurTime, 0),
		ReadingCnt: 1,
		Favour:     make([]dam.Favour, 0),
		FavourCnt:  0,
	}
	bookS.Reading = append(bookS.Reading, dam.DurTime{
		Start: readingStart,
		End:   "2199-12-05",
	})

	return dam.AddBook(bookS)
}

func AddFavour(id int, page, time, content, comment string) bool {
	favour := GetBook(id).Favour
	favour = append(favour, dam.Favour{
		Page:    page,
		Time:    time,
		Content: content,
		Comment: comment,
	})
	return dam.UpdateFavour(id, favour)
}

func SetTime(id, idInJson int, startOrEnd, time string) bool {
	reading := GetBook(id).Reading
	if len(reading) <= idInJson {
		return false
	}
	switch startOrEnd {
	case "start":
		reading[idInJson].Start = time
	case "end":
		reading[idInJson].End = time
	default:
		return false
	}
	dam.UpdateTime(id, reading)
	return true
}

func SetStartRead(id int, time string) bool {
	reading := GetBook(id).Reading
	reading = append(reading, dam.DurTime{
		Start: time,
		End:   "2199-12-05",
	})
	dam.UpdateTime(id, reading)
	dam.UpReadingCnt(id)
	return true
}

func Fix(id int, book, author, translator, publisher, tag string) bool {
	return dam.UpdateBookInfo(dam.Book{
		Id:         id,
		Book:       book,
		Author:     author,
		Translator: translator,
		Publisher:  publisher,
		Tag:        util.StringSplitBySpace(tag),
	})
}

func FixCover(id int, cover string) bool {
	return dam.UpdateBookCover(id, cover)
}

func FixFavour(id, idInJson int, page, time, content, comment string) bool {
	favour := GetBook(id).Favour
	if len(favour) <= idInJson {
		return false
	}
	favour[idInJson].Page = page
	favour[idInJson].Time = time
	favour[idInJson].Content = content
	favour[idInJson].Comment = comment

	return dam.UpdateFavour(id, favour)
}

func DelFavour(id, idInJson int) bool {
	favour := GetBook(id).Favour
	if len(favour) <= idInJson {
		return false
	}
	fav := append(favour[:idInJson], favour[idInJson+1:]...)

	return dam.UpdateFavour(id, fav)
}

func GetBookList() []dam.Book {
	return dam.GetBookList()
}

func GetBook(id int) dam.Book {
	return dam.GetBook(id)
}
