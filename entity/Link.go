package entity

import "time"

type Link struct {
	Id         int
	IdPage     int
	Url        string
	Visited    int
	CreatedAt  time.Time
	ModifiedAt time.Time
}
