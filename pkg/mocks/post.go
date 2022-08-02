package mocks

import (
	"time"

	"github.com/BradPreston/go-blog-backend/pkg/models"
)

var Posts = []models.Post{
	{
		ID:        1,
		Title:     "My first post",
		Image:     "test.png",
		Alt:       "missing image data",
		Body:      "This is my very first blog post",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
