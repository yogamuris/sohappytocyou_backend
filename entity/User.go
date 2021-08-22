package entity

import (
	"database/sql"
	"time"
)

type User struct {
	Id         int       `json:"id,omitempty"`
	Username   string    `json:"username,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	VerifiedAt sql.NullTime `json:"verified_at,omitempty"`
}