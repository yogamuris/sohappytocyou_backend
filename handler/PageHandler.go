package handler

import "net/http"

type PageHandler struct {

}

func NewPageHandler() PageHandler {
	return PageHandler{}
}

func (handler *PageHandler) Show(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}

func (handler *PageHandler) Create(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}

func (handler *PageHandler) Update(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}