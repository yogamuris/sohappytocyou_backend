package web

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type RegisterRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type VerifyRequest struct {
	Username string `json:"username,omitempty"`
	Token    string `json:"token,omitempty"`
}
