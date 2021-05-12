package db

import (
	"fmt"
)

const sqlBookshelf = `
create table book_%d
(
	id integer constraint book_%d_pk primary key,
	title text,
	title_origin text,
	author text,
	translator text,
	publisher text,
	cover text,
	tag text
);
INSERT INTO inc (name, val) VALUES ('book_%d', 0);
`
func createBookshelf(id int64) error {
	if !Connected {
		Connect()
	}
	err := Exec(fmt.Sprintf(sqlBookshelf, id, id, id)).Error
	return err
}

const sqlLike = `
create table like_%d
(
	id integer constraint like_%d_pk primary key autoincrement,
	origin text,
	time integer,
	content text,
	comment text,
	del integer default 0
);
`
func createLike(id int64) error {
	if !Connected {
		Connect()
	}
	err := Exec(fmt.Sprintf(sqlLike, id, id)).Error
	return err
}

const sqlBookRead = `
create table book_read_%d
(
	id integer constraint book_read_%d_pk primary key autoincrement,
	bid integer,
	start integer,
	end integer,
	del integer default 0
);
`
func createBookRead(id int64) error {
	if !Connected {
		Connect()
	}
	err := Exec(fmt.Sprintf(sqlBookRead, id, id)).Error
	return err
}

const sqlBookLike = `
create table book_like_%d
(
	id integer constraint book_like_%d_pk primary key autoincrement,
	bid integer,
	page integer,
	time integer,
	content text,
	comment text,
	del integer default 0
);
`
func createBookLike(id int64) error {
	if !Connected {
		Connect()
	}
	err := Exec(fmt.Sprintf(sqlBookLike, id, id)).Error
	return err
}

func BidValid(uuid, bid int64) bool {
	if !Connected {
		Connect()
	}
	return bid >= 1 && bid <= GetInc(fmt.Sprintf("book_%d", uuid))
}