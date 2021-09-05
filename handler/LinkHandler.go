package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yogamuris/sohappytocyou/entity/web"
	"github.com/yogamuris/sohappytocyou/service"
	"log"
	"net/http"
	"strconv"
)

type LinkHandler struct {
	Service service.LinkService
}

func NewLinkHandler(service service.LinkService) LinkHandler {
	return LinkHandler{Service: service}
}

func (handler *LinkHandler) List(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	username := params["username"]

	if username != request.Context().Value("username") {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	linkResponse, err := handler.Service.List(request.Context(), username)
	encoder := json.NewEncoder(writer)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Data: linkResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = encoder.Encode(webResponse)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (handler *LinkHandler) Show(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	username := params["username"]
	if username != request.Context().Value("username") {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	idLink := params["id"]
	id, _ := strconv.Atoi(idLink)

	linkResponse, err := handler.Service.Show(request.Context(), id)
	encoder := json.NewEncoder(writer)

	if err != nil {
		if err.Error() == "link not found" {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Data: linkResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = encoder.Encode(webResponse)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (handler *LinkHandler) Create(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	username := params["username"]
	if username != request.Context().Value("username") {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	linkSaveRequest := web.LinkSaveRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&linkSaveRequest)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	linkResponse, err := handler.Service.Save(request.Context(), linkSaveRequest)
	encoder := json.NewEncoder(writer)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Data: linkResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = encoder.Encode(webResponse)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (handler *LinkHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	username := params["username"]
	if username != request.Context().Value("username") {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	linkDeleteRequest := web.LinkDeleteRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&linkDeleteRequest)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	ok := handler.Service.Delete(request.Context(), linkDeleteRequest)
	if !ok {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)

}
