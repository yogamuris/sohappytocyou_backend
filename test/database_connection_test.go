package test

import (
	"github.com/yogamuris/sohappytocyou/database"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	db := database.GetTestDB()

	err := db.Ping()
	if err != nil {
		t.Error("Database message")
	}
}