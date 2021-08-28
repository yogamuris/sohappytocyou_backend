package entity

import (
	"database/sql"
	"time"
)

type User struct {
	Id         int
	Username   string
	Email      string
	Password   string
	CreatedAt  time.Time
	VerifiedAt sql.NullTime
}
