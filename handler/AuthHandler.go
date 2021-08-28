package handler

import (
	"encoding/json"
	"github.com/yogamuris/sohappytocyou/entity/web"
	"github.com/yogamuris/sohappytocyou/service"
	"net/http"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return AuthHandler{
		AuthService: service,
	}
}

func (handler *AuthHandler) Login(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}

func (handler *AuthHandler) Register(writer http.ResponseWriter, request *http.Request) {
	registerRequest := web.RegisterRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&registerRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userResponse, err := handler.AuthService.Register(request.Context(), registerRequest)

	encoder := json.NewEncoder(writer)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		if err.Error() == "unique field message" {
			encoder.Encode(web.WebResponse{
				Code: http.StatusInternalServerError,
				Data: "Username / Email tidak tersedia",
			})
		}

		return
	}

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Data: userResponse.Message,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = encoder.Encode(webResponse)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (handler *AuthHandler) Verify(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}
