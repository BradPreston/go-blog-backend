package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port string = ":8080"

func main() {
	router := mux.NewRouter()

	log.Println("API is running on port:", port)
	http.ListenAndServe(port, router)
}