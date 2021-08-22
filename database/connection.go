package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Println(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(20 * time.Minute)

	return db
}