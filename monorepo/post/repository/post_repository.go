package repository

import (
	"database/sql"
	"errors"
	"v1/monorepo/post/model"
)

type PostRepository interface {
	FindPost() ([]model.Post, error)
	FindPostByName() ([]model.Post, error)
	Create(userId uint64, postBody model.PostDTO) error
	Update() error
	Delete() error
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}

func (p *postRepository) FindPost() ([]model.Post, error) {
	return []model.Post{}, errors.New("hello")
}

func (p *postRepository) FindPostByName() ([]model.Post, error) {
	return []model.Post{}, errors.New("hello")
}

func (p *postRepository) Create(userId uint64, postBody model.PostDTO) error {
	return nil
}

func (p *postRepository) Update() error {
	return nil
}

func (p *postRepository) Delete() error {
	return nil
}
