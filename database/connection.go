package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yogamuris/sohappytocyou/helper"
	"log"
	"time"
)

func GetTestDB() *sql.DB {
	dbHost := helper.GetEnv("../.env", "TEST_DB_HOST")
	dbDriver := helper.GetEnv("../.env", "TEST_DB_DRIVER")
	dbUser := helper.GetEnv("../.env", "TEST_DB_USER")
	dbPassword := helper.GetEnv("../.env", "TEST_DB_PASSWORD")
	dbName := helper.GetEnv("../.env", "TEST_DB_NAME")
	dbPort := helper.GetEnv("../.env", "TEST_DB_PORT")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open(dbDriver, dataSource)
	if err != nil {
		log.Println(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(20 * time.Minute)

	return db
}

func GetDB() *sql.DB {
	dbHost := helper.GetEnv(".env", "DB_HOST")
	dbDriver := helper.GetEnv(".env", "DB_DRIVER")
	dbUser := helper.GetEnv(".env", "DB_USER")
	dbPassword := helper.GetEnv(".env", "DB_PASSWORD")
	dbName := helper.GetEnv(".env", "DB_NAME")
	dbPort := helper.GetEnv(".env", "DB_PORT")



	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open(dbDriver, dataSource)
	if err != nil {
		log.Println(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(20 * time.Minute)

	return db
}