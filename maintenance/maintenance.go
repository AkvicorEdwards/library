package maintenance

import (
	"library/dam"
	"log"
	"os"
)

const sql = `
create table inc
(
    name text,
    val  integer
);
` + `
create table user
(
    uuid     integer
        constraint user_pk
            primary key,
    username text,
    nickname text,
    password text
);

create unique index user_username_uindex
    on user (username);
` + `
create table books
(
    id        integer
        constraint books_pk
            primary key,
    book        text,
    author      text,
    translator  text,
    publisher   text,
    cover       text,
    tag         text,
    reading     text,
    reading_cnt integer default 1,
    favour      text,
    favour_cnt  integer default 0
);
` + `
INSERT INTO inc (name, val) VALUES ('user', 0);
INSERT INTO inc (name, val) VALUES ('books', 1);
INSERT INTO books (id, book, author, translator, publisher, 
	cover, tag, reading, reading_cnt, favour, favour_cnt) VALUES
	(1, 'Complex', '', '', '', '', '[""]', '[{"end_time": "2199-12-05",
	"start_time": "2020-03-19"}]', 1, '[]', 0);
`

func InitDatabase() {
	if !IsFile("library.db") {
		log.Println("library.db do not exist")
		os.Exit(-1)
	}

	err := dam.Exec(sql).Error
	if err != nil {
		log.Println(err)
		os.Exit(-2)
	}

	log.Println("Finished")

	os.Exit(0)
}

func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	if !Exists(path) {
		return false
	}
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}