package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/BradPreston/go-blog-backend/internal/mocks"
	"github.com/BradPreston/go-blog-backend/internal/models"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var post models.Post

	json.Unmarshal(body, &post)

	post.ID = rand.Intn(100)
	mocks.Posts = append(mocks.Posts, post)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Created")
}
