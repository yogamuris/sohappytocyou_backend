package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/yogamuris/sohappytocyou/database"
	"github.com/yogamuris/sohappytocyou/handler"
	"github.com/yogamuris/sohappytocyou/repository"
	"github.com/yogamuris/sohappytocyou/service"
	"log"
	"net/http"
)

func main() {
	db := database.GetDB()
	defer db.Close()

	validate := validator.New()
	authHandler := handler.NewAuthHandler()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userHandler := handler.NewUserHandler(userService)


	router := mux.NewRouter()
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/register", userHandler.Create).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented yet"))
}