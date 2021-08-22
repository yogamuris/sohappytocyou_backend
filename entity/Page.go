package entity

import "image/color"

type Page struct {
	Id          int         `json:"id,omitempty"`
	Username    string      `json:"username,omitempty" json:"username,omitempty"`
	Background  color.Color `json:"background,omitempty" json:"background,omitempty"`
	Photo       string      `json:"photo,omitempty" json:"photo,omitempty"`
	Description string      `json:"description,omitempty" json:"description,omitempty"`
	Links       []Link      `json:"links,omitempty" json:"links,omitempty"`
}
