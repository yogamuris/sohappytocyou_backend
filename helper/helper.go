package helper

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		if errorRollback != nil {
			log.Println(errorRollback)
		}
		log.Println(err)
	} else {
		errorCommit := tx.Commit()
		if errorCommit != nil {
			log.Println(errorCommit)
		}
	}
}

func GetEnv(path, key string) string {
	err := godotenv.Load(path)
	if err != nil {
		log.Println(err)
	}

	return os.Getenv(key)
}
