package entity

type Link struct {
	Id      int    `json:"id,omitempty"`
	Url     string `json:"url,omitempty"`
	Visited int  `json:"visited,omitempty"`
}