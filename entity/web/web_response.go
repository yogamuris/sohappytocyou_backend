package web

type WebResponse struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
}