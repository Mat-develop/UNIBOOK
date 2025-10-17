package service

import (
	"errors"
	"v1/monorepo/post/model"
	"v1/monorepo/post/repository"
)

type PostService interface {
	GetPost(accId uint) []model.Post
	GetPostByName() []model.Post
	CreatePost(userId uint64, postBody model.PostDTO) error
	UpdatePost()
	DeletePost()
}

type postService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return &postService{postRepository: postRepository}
}

func (p *postService) GetPost(accId uint) []model.Post {
	return []model.Post{}
}

func (p *postService) CreatePost(userId uint64, postBody model.PostDTO) error {
	return errors.New("hello there")
}

func (p *postService) GetPostByName() []model.Post {
	return []model.Post{}
}

func (p *postService) UpdatePost() {}

func (p *postService) DeletePost() {}
