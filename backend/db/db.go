package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB_PATH string
var sdb *sql.DB = nil

func Init(path string) {
	DB_PATH = path
	db, err := sql.Open("sqlite3", DB_PATH)
	if err != nil {
		log.Fatalln("Error:", err)
	}
    sdb = db

	sqls := []string{
		"CREATE TABLE IF NOT EXISTS `users` (`uid` INTEGER PRIMARY KEY AUTOINCREMENT, `name` VARCHAR(64) NOT NULL UNIQUE, `passwd` VARCHAR(64) NOT NULL)",
	}

	for _, s := range sqls {
		if stmt, err := sdb.Prepare(s); err != nil {
			log.Fatalln("Error:", err)
		} else {
			if _, err := stmt.Exec(); err != nil {
				log.Fatalln("Error:", err)
			}
		}
	}
}

func Uninit () {
    if sdb != nil {
        sdb.Close()
    }
}
