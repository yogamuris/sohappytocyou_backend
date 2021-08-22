package handler

import "net/http"

type AuthHandler struct {

}

func NewAuthHandler() AuthHandler {
	return AuthHandler{}
}

func (handler *AuthHandler) Login(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}

func (handler *AuthHandler) Register(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}

func (handler *AuthHandler) Verify(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}