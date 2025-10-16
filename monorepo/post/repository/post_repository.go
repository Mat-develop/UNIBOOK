package repository

import (
	"database/sql"
)

type PostRepository interface {
	// FindPost() []model.Post
	// FindPostByName() []model.Post
	// Create()
	// Update()
	// Delete()
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}
