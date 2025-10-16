package handlers

import (
	"net/http"
	"v1/monorepo/post/service"
)

type PostHandler interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	GetPostByTittle(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
}

type postHandler struct {
	service service.PostService
}

func NewPostHandler(handlerService service.PostService) PostHandler {
	return &postHandler{service: handlerService}
}

func (p *postHandler) GetPosts(w http.ResponseWriter, r *http.Request)        {}
func (p *postHandler) GetPostByTittle(w http.ResponseWriter, r *http.Request) {}
func (p *postHandler) CreatePost(w http.ResponseWriter, r *http.Request)      {}
func (p *postHandler) UpdatePost(w http.ResponseWriter, r *http.Request)      {}
func (p *postHandler) DeletePost(w http.ResponseWriter, r *http.Request)      {}
