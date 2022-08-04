package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BradPreston/go-blog-backend/internal/config"
	"github.com/BradPreston/go-blog-backend/internal/driver"
	"github.com/joho/godotenv"
)

var port string = ":8080"

var app config.AppConfig

func main() {
	env, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connecting to MongoDB")
	postsCollection, err := driver.ConnectMongoDB(env["MONGODB_URI"], "blog", "posts")
	if err != nil {
		log.Fatal("MongoDB connection failed...")
	}
	fmt.Println("Connected to MongoDB")
	fmt.Println(postsCollection)

	srv := &http.Server{
		Addr:         port,
		Handler:      Routes(&app),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Println("API is running on port:", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
