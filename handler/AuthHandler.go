package handler

import "net/http"

type AuthHandler struct {

}

func NewAuthHandler() AuthHandler {
	return AuthHandler{}
}

func (handler *AuthHandler) Login(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Not Implemented"))
}