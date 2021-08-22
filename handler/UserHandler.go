package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yogamuris/sohappytocyou/entity/web"
	"github.com/yogamuris/sohappytocyou/service"
	"net/http"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{
		UserService: userService,
	}
}

func (handler *UserHandler) Create(writer http.ResponseWriter, request *http.Request) {
	userCreateRequest := web.UserCreateRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(userCreateRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userResponse, err := handler.UserService.Create(request.Context(), userCreateRequest)

	encoder := json.NewEncoder(writer)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeOkRequest(writer, encoder, userResponse)
}

func (handler *UserHandler) ChangePassword(writer http.ResponseWriter, request *http.Request) {
	userChangePasswordRequest := web.UserChangePasswordRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userChangePasswordRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(writer)

	userResponse, err := handler.UserService.ChangePassword(request.Context(), userChangePasswordRequest)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeOkRequest(writer, encoder, userResponse)
}

func (handler *UserHandler) FindByUsername(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	username := params["username"]

	encoder := json.NewEncoder(writer)
	userResponse, err := handler.UserService.FindByUsername(request.Context(), username)
	if err != nil {
		if err.Error() == "user not found" {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeOkRequest(writer, encoder, userResponse)
}

func writeOkRequest(writer http.ResponseWriter, encoder *json.Encoder, userResponse web.UserResponse) {
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Data: userResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := encoder.Encode(webResponse)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}