package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func Connect() {
	dsn := "root:Anshit123@#@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=True"
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func GetDB() *sql.DB {
	return db
}
