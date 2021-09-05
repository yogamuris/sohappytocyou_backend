package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yogamuris/sohappytocyou/entity/web"
	"log"

	//"github.com/yogamuris/sohappytocyou/service"
	"github.com/yogamuris/sohappytocyou/service"
	"net/http"
)

type PageHandler struct {
	PageService service.PageService
}

func NewPageHandler(pageService service.PageService) PageHandler {
	return PageHandler{PageService: pageService}
}

func (handler *PageHandler) Show(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	username := params["username"]

	pageResponse, err := handler.PageService.Show(request.Context(), username)
	encoder := json.NewEncoder(writer)

	if err != nil {
		if err.Error() == "page not found" {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Data: pageResponse,
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

func (handler *PageHandler) Create(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	username := params["username"]
	if username != request.Context().Value("username") {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	pageSaveRequest := web.PageSaveRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&pageSaveRequest)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	pageResponse, err := handler.PageService.Save(request.Context(), pageSaveRequest)

	encoder := json.NewEncoder(writer)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Data: pageResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = encoder.Encode(webResponse)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (handler *PageHandler) Update(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	username := params["username"]
	if username != request.Context().Value("username") {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	pageUpdateRequest := web.PageUpdateRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&pageUpdateRequest)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	pageResponse, err := handler.PageService.Update(request.Context(), pageUpdateRequest)
	encoder := json.NewEncoder(writer)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Data: pageResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = encoder.Encode(webResponse)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

}
