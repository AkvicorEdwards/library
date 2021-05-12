package db

import (
	"gorm.io/gorm/logger"
	"log"
	"library/def"
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Inc struct {
	Name string // table name
	Val int64
}

var (
	db        *gorm.DB
	Connected = false
	lockUser   = sync.RWMutex{}
)

func Connect() {
	if Connected {
		return
	}
	var err error
	db, err = gorm.Open(sqlite.Open(def.DBFilename), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Println(err)
		return
	}
	Connected = true
}

func Exec(sql string, values ...interface{}) *gorm.DB {
	if !Connected {
		Connect()
	}
	return db.Exec(sql, values...)
}

func Raw(sql string, values ...interface{}) *gorm.DB {
	if !Connected {
		Connect()
	}
	return db.Raw(sql, values...)
}

func GetInc(name string) int64 {
	if !Connected {
		Connect()
	}
	ic := Inc{}
	db.Table("inc").Where("name=?", name).First(&ic)
	return ic.Val
}

func updateInc(name string, val int64) bool {
	if !Connected {
		Connect()
	}
	return db.Table("inc").Where("name=?", name).Update("val", val).Error == nil
}

func Lock() {
	lockUser.Lock()
}

func UnixTime() int64 {
	return time.Now().Unix()
}

func Abs(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}
