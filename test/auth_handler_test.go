package test

import (
	"bytes"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/yogamuris/sohappytocyou/database"
	"github.com/yogamuris/sohappytocyou/handler"
	"github.com/yogamuris/sohappytocyou/repository"
	"github.com/yogamuris/sohappytocyou/service"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	db = database.GetTestDB()
	validate = validator.New()

	userRepository = repository.NewUserRepository()
	authService = service.NewAuthService(userRepository, db, validate)
	authHandler = handler.NewAuthHandler(authService)
)

func TestRegister(t *testing.T) {
	tests := []struct {
		name string
		url string
		method string
		body []byte
		statusCode int
	}{
		{
			name: "test success register",
			url: "http://localhost:8080/auth/register",
			method: "POST",
			body: []byte(`{"username": "adhiyoga","email": "test@test.com","password": "newpassword"}`),
			statusCode: http.StatusOK,
		},
		{
			name: "test registration but username exists",
			url: "http://localhost:8080/auth/register",
			method: "POST",
			body: []byte(`{"username": "adhiyoga","email": "test2@test.com","password": "newpassword"}`),
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "test registration but email exists",
			url: "http://localhost:8080/auth/register",
			method: "POST",
			body: []byte(`{"username": "adhiyoga2","email": "test@test.com","password": "newpassword"}`),
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "test register without username",
			url: "http://localhost:8080/auth/register",
			method: "POST",
			body: []byte(`{"email": "test32@test.com","password": "newpassword"}`),
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "test register without email",
			url: "http://localhost:8080/auth/register",
			method: "POST",
			body: []byte(`{"username": "adhiyoga31", "password": "newpassword3"}`),
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "test register without password",
			url: "http://localhost:8080/auth/register",
			method: "POST",
			body: []byte(`{"username": "adhiyoga23", "email": "test34@test.com"}`),
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		log.Println(tt.name)
		r := httptest.NewRequest(tt.method, tt.url, bytes.NewBuffer(tt.body))
		w := httptest.NewRecorder()
		authHandler.Register(w, r)
		resp := w.Result()
		assert.Equal(t, tt.statusCode, resp.StatusCode)
	}
}

func TestLogin(t *testing.T) {
	test := []struct{
		name string
		url string
		method string
		body []byte
		statusCode int
	}{
		{
			name: "Test success login",
			url: "http://localhost:8080/auth/login",
			method: "POST",
			body: []byte(`{"username":"adhiyoga", "password": "newpassword"}`),
			statusCode: http.StatusOK,
		},
		{
			name: "Test wrong password",
			url: "http://localhost:8080/auth/login",
			method: "POST",
			body: []byte(`{"username":"adhiyoga", "password": "newpassword2"}`),
			statusCode: http.StatusUnauthorized,
		},
		{
			name: "Test unregister user",
			url: "http://localhost:8080/auth/login",
			method: "POST",
			body: []byte(`{"username":"adhiyoga2", "password": "newpassword"}`),
			statusCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range test {
		log.Println(tt.name)
		r := httptest.NewRequest(tt.method, tt.url, bytes.NewBuffer(tt.body))
		w := httptest.NewRecorder()
		authHandler.Login(w, r)

		resp := w.Result()
		assert.Equal(t, tt.statusCode, resp.StatusCode)
	}
}