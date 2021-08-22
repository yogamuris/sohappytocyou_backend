package web

type UserChangePasswordRequest struct {
	Id       int
	Password string `json:"password,omitempty"`
}
