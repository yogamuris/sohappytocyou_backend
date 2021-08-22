package handler

import (
	"encoding/json"
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

func (handler UserHandler) Create(writer http.ResponseWriter, request *http.Request) {
	userCreateRequest := web.UserCreateRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userCreateRequest)
	if err != nil {
		panic(err)
	}

	userResponse, err := handler.UserService.Create(request.Context(), userCreateRequest)
	if err != nil {
		panic(err)
	}

	webResponse := web.WebResponse{
		Code: 200,
		Data: userResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}