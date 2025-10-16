package service

import "v1/monorepo/post/repository"

type PostService interface {
	// GetPost(accId uint) []model.Post
	// GetPostByName() []model.Post
	// CreatePost()
	// UpdatePost()
	// DeletePost()
}

type postService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return &postService{postRepository: postRepository}
}
