package web

import "github.com/dgrijalva/jwt-go"

type AuthResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Token struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
