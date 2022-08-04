package handlers

import "github.com/BradPreston/go-blog-backend/internal/config"

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
