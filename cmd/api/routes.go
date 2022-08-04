package main

import (
	"net/http"

	"github.com/BradPreston/go-blog-backend/internal/config"
	"github.com/BradPreston/go-blog-backend/internal/handlers"
	"github.com/gorilla/mux"
)

func Routes(app *config.AppConfig) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/posts", handlers.GetAllPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts", handlers.AddPost).Methods(http.MethodPost)
	router.HandleFunc("/posts/{id}", handlers.GetPost).Methods(http.MethodGet)
	router.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods(http.MethodPut)
	router.HandleFunc("/posts/{id}", handlers.DeletePost).Methods(http.MethodDelete)

	return router
}
