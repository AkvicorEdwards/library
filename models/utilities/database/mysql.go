package database

import (
	"database/sql"
	"fmt"
	"library/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *MySQL

func Init() {
	DB = NewByConfig()
}

type SQL interface {
	SetDSN(user string, password string, database string, charset string)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (rows *sql.Rows, err error)
}

type MySQL struct {
	DataSourceName string
	DB *sql.DB
}

func NewByConfig() (mysql *MySQL) {
	mysql = &MySQL{}
	mysql.SetDSN(config.Data.Mysql.User, config.Data.Mysql.Password,
				 config.Data.Mysql.Database, config.Data.Mysql.Charset)
	return
}

func NewByArgs(user string, password string, database string, charset string) (mysql *MySQL) {
	mysql = &MySQL{}
	mysql.SetDSN(user, password, database, charset)
	return
}

func New() *MySQL {
	return &MySQL{}
}

func (c *MySQL)SetDSN(user string, password string, database string, charset string) {
	c.DataSourceName = fmt.Sprintf("%s:%s@/%s?charset=%s", user, password, database, charset)
}

func (c *MySQL) Exec(query string, args ...interface{}) (result sql.Result, err error) {
	c.DB, err = sql.Open("mysql", c.DataSourceName)
	if err != nil {
		return
	}
	defer func() {
		_ = c.DB.Close()
	}()

	result, err = c.DB.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (c *MySQL) Query(query string, args ...interface{}) (rows *sql.Rows, err error) {
	c.DB, err = sql.Open("mysql", c.DataSourceName)
	if err != nil {
		return
	}
	defer func() {
		_ = c.DB.Close()
	}()
	rows, err = c.DB.Query(query, args...)
	return
}
