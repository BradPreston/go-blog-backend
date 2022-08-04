package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/BradPreston/go-blog-backend/internal/mocks"
	"github.com/BradPreston/go-blog-backend/internal/models"
	"github.com/gorilla/mux"
)

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var updatedPost models.Post

	json.Unmarshal(body, &updatedPost)

	for index, post := range mocks.Posts {
		if post.ID == id {
			post.Title = updatedPost.Title
			post.Alt = updatedPost.Alt
			post.Image = updatedPost.Image
			post.Body = updatedPost.Body
			post.CreatedAt = time.Now()
			post.UpdatedAt = time.Now()

			mocks.Posts[index] = post
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("updated")
			break
		}
	}
}
