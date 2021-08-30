package web

type LinkListRequest struct {
	IdPage int `json:"id_page,omitempty"`
}

type LinkSaveRequest struct {
	IdPage int    `json:"id_page,omitempty"`
	Url    string `json:"url,omitempty"`
}

type LinkGetRequest struct {
	Id int `json:"id,omitempty"`
}

type LinkUpdateRequest struct {
	Id  int    `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
}

type LinkDeleteRequest struct {
	Id int `json:"id,omitempty"`
}
