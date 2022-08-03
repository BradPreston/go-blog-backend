package main

import (
	"log"
	"net/http"
	"time"
)

var port string = ":8080"

type Config struct {
	port        string
	environment string
}

type application struct {
	config Config
}

func main() {
	var cfg Config
	cfg.port = ":8080"
	cfg.environment = "development"

	app := &application{
		config: cfg,
	}

	srv := &http.Server{
		Addr:         cfg.port,
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Println("API is running on port:", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
