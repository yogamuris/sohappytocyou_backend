package entity

import "time"

type Page struct {
	Id          int
	IdUser      int
	Username    string
	Background  string
	Photo       string
	Description string
	CreatedAt   time.Time
	ModifiedAt  time.Time
	Links       []Link
}
