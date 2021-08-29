package handler

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/yogamuris/sohappytocyou/entity/web"
	"github.com/yogamuris/sohappytocyou/service"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return AuthHandler{
		AuthService: service,
	}
}

func GenerateJWT(username string, expiration time.Time) (string, error) {
	var jwtKey = []byte("lets_groove_tonight")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["exp"] = expiration.Unix()

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (handler *AuthHandler) Login(writer http.ResponseWriter, request *http.Request) {
	var loginRequest web.LoginRequest
	err := json.NewDecoder(request.Body).Decode(&loginRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := handler.AuthService.Login(request.Context(), loginRequest)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	checkPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)) != nil

	if !checkPassword {
		log.Println("salah pwd")
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	expiration := time.Now().Add(1 * time.Minute)

	validToken, err := GenerateJWT(user.Username, expiration)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	var token web.Token
	token.Username = user.Username
	token.Token = validToken

	http.SetCookie(writer, &http.Cookie{
		Name:    "sohappytocyou_token",
		Value:   validToken,
		Expires: expiration,
		Path:    "/",
	})

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(token)
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
