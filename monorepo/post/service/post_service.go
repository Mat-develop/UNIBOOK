package service

import "v1/monorepo/post/model"

type PostService interface {
	GetPost(accId uint) []model.Post
	GetPostByName() []model.Post
	CreatePost()
	UpdatePost()
	DeletePost()
}

type postService struct {
	postRepo repository
}

func NewPostService(postRepo repository) PostService {

}
