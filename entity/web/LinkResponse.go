package web

type LinkResponse struct {
	Id      int    `json:"id"`
	IdPage  int    `json:"id_page"`
	Url     string `json:"url"`
	Visited int    `json:"visited"`
}

type LinkListResponse struct {
	Links []LinkResponse
}
