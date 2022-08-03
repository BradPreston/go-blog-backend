package main

import (
	"log"
	"net/http"

	"github.com/BradPreston/go-blog-backend/pkg/handlers"
	"github.com/gorilla/mux"
)

var port string = ":8080"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/posts", handlers.GetAllPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts", handlers.AddPost).Methods(http.MethodPost)
	router.HandleFunc("/posts/{id}", handlers.GetPost).Methods(http.MethodGet)
	router.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods(http.MethodPut)

	log.Println("API is running on port:", port)
	http.ListenAndServe(port, router)
}
