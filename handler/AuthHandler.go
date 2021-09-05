package handler

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/yogamuris/sohappytocyou/entity/web"
	"github.com/yogamuris/sohappytocyou/helper"
	"github.com/yogamuris/sohappytocyou/service"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type AuthHandler struct {
	AuthService service.AuthService
}

var jwtKey = []byte(helper.GetEnv(".env", "JWT_KEY"))

func NewAuthHandler(service service.AuthService) AuthHandler {
	return AuthHandler{
		AuthService: service,
	}
}

func GenerateJWT(username string, expiration time.Time) (string, error) {

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

	checkPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))

	if checkPassword != nil {
		log.Println("wrong password")
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	expiration := time.Now().Add(60 * time.Minute)

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

func (handler *AuthHandler) Refresh(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("sohappytocyou_token")
	if err != nil {
		if err == http.ErrNoCookie {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	token := cookie.Value
	claims := web.Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 5*time.Minute {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString(jwtKey)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(writer, &http.Cookie{
		Name:    "sohappytocyou_token",
		Value:   tokenString,
		Expires: expirationTime,
		Path:    "/",
	})
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
