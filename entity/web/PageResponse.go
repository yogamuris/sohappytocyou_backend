package web

import (
	"github.com/yogamuris/sohappytocyou/entity"
)

type PageResponse struct {
	Id          int           `json:"id"`
	Username    string        `json:"username"`
	Background  string        `json:"background"`
	Photo       string        `json:"photo"`
	Description string        `json:"description"`
	Links       []entity.Link `json:"links"`
}
