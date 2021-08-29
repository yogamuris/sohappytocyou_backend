package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/yogamuris/sohappytocyou/database"
	"github.com/yogamuris/sohappytocyou/handler"
	"github.com/yogamuris/sohappytocyou/middleware"
	"github.com/yogamuris/sohappytocyou/repository"
	"github.com/yogamuris/sohappytocyou/service"
	"log"
	"net/http"
)

func main() {
	db := database.GetDB()
	defer db.Close()

	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userHandler := handler.NewUserHandler(userService)

	authService := service.NewAuthServie(userRepository, db, validate)
	authHandler := handler.NewAuthHandler(authService)

	pageRepository := repository.NewPageRepository()
	pageService := service.NewPageService(pageRepository, db, validate)
	pageHandler := handler.NewPageHandler(pageService)

	linkRepository := repository.NewLinkRepository()
	linkService := service.NewLinkService(linkRepository, db, validate)
	linkHandler := handler.NewLinkHandler(linkService)

	router := mux.NewRouter()
	router.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	//router.HandleFunc("/auth/verify", authHandler.Verify).Methods("POST")

	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.Use(middleware.AuthMiddleware)

	userRouter.HandleFunc("/{username}", userHandler.FindByUsername).Methods("GET")
	userRouter.HandleFunc("/{username}/change-password", userHandler.ChangePassword).Methods("PUT")

	userRouter.HandleFunc("/{username}/page", pageHandler.Show).Methods("GET")
	userRouter.HandleFunc("/{username}/page", pageHandler.Create).Methods("POST")
	userRouter.HandleFunc("/{username}/page/update", pageHandler.Update).Methods("PUT")

	userRouter.HandleFunc("/{username}/page/links", linkHandler.List).Methods("GET")
	userRouter.HandleFunc("/{username}/page/links", linkHandler.Create).Methods("POST")
	userRouter.HandleFunc("/{username}/page/links/{id}", linkHandler.Show).Methods("GET")
	userRouter.HandleFunc("/{username}/page/links/{id}/delete", linkHandler.Delete).Methods("DELETE")

	//router.HandleFunc("/user/{username}/page/analytic", NotImplemented).Methods("GET")
	//router.HandleFunc("/user/{username}/page/links/analytic", NotImplemented).Methods("GET")
	//router.HandleFunc("/user/{username}/page/links/{id}/analytic", NotImplemented).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented yet"))
}
