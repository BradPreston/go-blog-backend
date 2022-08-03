package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/BradPreston/go-blog-backend/pkg/models"
)

func TestAddPost(t *testing.T) {
	var post models.Post

	post.ID = 1
	post.Title = "test post"
	post.Image = "test.png"
	post.Alt = "test image"
	post.Body = "test body"
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	data, err := json.Marshal(post)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddPost)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned the wrong status code: got %v but expected %v", status, http.StatusOK)
	}
}
