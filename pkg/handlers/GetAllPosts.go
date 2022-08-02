package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/BradPreston/go-blog-backend/pkg/mocks"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Posts)
}
