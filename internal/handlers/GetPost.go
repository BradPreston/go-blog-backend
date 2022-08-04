package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BradPreston/go-blog-backend/internal/mocks"
	"github.com/gorilla/mux"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, post := range mocks.Posts {
		if post.ID == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(post)
			break
		}
	}
}
