package test

import (
	"github.com/yogamuris/sohappytocyou/database"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	db := database.GetDB()

	err := db.Ping()
	if err != nil {
		t.Error("Database error")
	}
}