package web

type PageRequest struct {
	Username string `json:"username,omitempty"`
}

type PageSaveRequest struct {
	Username    string `json:"username,omitempty"`
	Background  string `json:"background,omitempty"`
	Photo       string `json:"photo,omitempty"`
	Description string `json:"description,omitempty"`
}

type PageUpdateRequest struct {
	Username    string `json:"username,omitempty"`
	Background  string `json:"background,omitempty"`
	Photo       string `json:"photo,omitempty"`
	Description string `json:"description,omitempty"`
}
