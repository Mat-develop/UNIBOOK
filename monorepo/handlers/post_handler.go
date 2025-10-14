package handlers

import (
	"net/http"
	"v1/monorepo/users/service"
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

}

func GetPosts(w http.ResponseWriter, r *http.Request)        {}
func GetPostByTittle(w http.ResponseWriter, r *http.Request) {}
func CreatePost(w http.ResponseWriter, r *http.Request)      {}
func UpdatePost(w http.ResponseWriter, r *http.Request)      {}
func DeletePost(w http.ResponseWriter, r *http.Request)      {}
