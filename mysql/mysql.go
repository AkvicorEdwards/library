package mysql

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DEFAULT *MySQL

type MySQL struct {
	User string
	Password string
	Host string
	DBName string
	Charset string
	ParseTime string
	Loc string
}

func InitBySelf() {
	DEFAULT = &MySQL{
		User:      "root",
		Password:  "password",
		Host:      "localhost",
		DBName:    "book",
		Charset:   "utf8mb4",
		ParseTime: "true",
		Loc:       "Local",
	}
}


func SetDEFAULT(c *MySQL) {
	DEFAULT = c
}



func Test() {
	var result []TableUser

	_, err := Row(&result, "SELECT * FROM users WHERE id > 1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func Open(m *MySQL) (*gorm.DB, error) {
	return gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=%s&loc=%s",
		m.User,
		m.Password,
		m.Host,
		m.DBName,
		m.Charset,
		m.ParseTime,
		m.Loc,
		))
}

func Close(db *gorm.DB) {
	_ = db.Close()
}

func Exec(sql string, values ...interface{}) (*gorm.DB, error) {
	db, err := Open(DEFAULT)
	if err != nil {
		return nil, err
	}
	defer Close(db)
	gdb := db.Exec(sql, values...)
	return gdb, gdb.Error
}

func Row(st interface{}, sql string, values ...interface{}) (*gorm.DB, error) {
	db, err := Open(DEFAULT)
	if err != nil {
		return nil, err
	}
	defer Close(db)
	gdb := db.Raw(sql, values...).Scan(st)
	return gdb, gdb.Error
}

func Query(sql string, values ...interface{}) (*sql.Rows, error) {
	db, err := Open(DEFAULT)
	if err != nil {
		return nil, err
	}
	defer Close(db)
	gdb := db.Raw(sql, values...)
	gdbs, err := gdb.Rows()

	return gdbs, nil
}