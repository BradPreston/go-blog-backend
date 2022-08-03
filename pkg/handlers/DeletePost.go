package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BradPreston/go-blog-backend/pkg/mocks"
	"github.com/gorilla/mux"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, post := range mocks.Posts {
		if post.ID == id {
			mocks.Posts = append(mocks.Posts[:index], mocks.Posts[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}
