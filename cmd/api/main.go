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

	log.Println("API is running on port:", port)
	http.ListenAndServe(port, router)
}
