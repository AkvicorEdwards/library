package repair

import (
	"library/db"
	"library/def"
	"log"
	"os"

	"github.com/AkvicorEdwards/util"
)

const sql = `
create table inc
(
    name text,
    val  integer
);

create table user
(
	uuid integer constraint user_pk primary key,
	username text,
	password text,
	profile_photo text,
	nickname text,
	last_login integer default 0
);

create unique index user_username_uindex on user (username);

INSERT INTO inc (name, val) VALUES ('user', 0);
`


func InitDatabase() {
	stat := util.FileStat(def.DBFilename)
	if stat == 0 || stat == 2 {
		if stat == 2 {
			err := os.Remove(def.DBFilename)
			if err != nil {
				log.Println("Cannot delete file")
				os.Exit(-4)
			}
		}
		f, err := os.Create(def.DBFilename)
		if err != nil {
			log.Println("Cannot create file")
			os.Exit(-1)
		}
		err = f.Close()
		if err != nil {
			log.Println("Create failed")
			os.Exit(-2)
		}
	} else if stat == 1 {
		log.Printf("%s is a directory\n", def.DBFilename)
		os.Exit(-3)
	}

	err := db.Exec(sql).Error
	if err != nil {
		log.Println(err)
		os.Exit(-5)
	}

	log.Println("Init DB Finished")
}