package handler

import "net/http"

type LinkHandler struct {

}

func NewLinkHandler() LinkHandler {
	return LinkHandler{}
}

func (handler *LinkHandler) List(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}

func (handler *LinkHandler) Show(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}

func (handler *LinkHandler) Create(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}

func (handler *LinkHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}