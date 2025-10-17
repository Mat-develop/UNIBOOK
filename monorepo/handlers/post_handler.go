package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"v1/monorepo/post/model"
	"v1/monorepo/post/service"
	"v1/monorepo/util/authentication"
	"v1/monorepo/util/response"
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

func (p *postHandler) GetPosts(w http.ResponseWriter, r *http.Request) {

}

func (p *postHandler) GetPostByTittle(w http.ResponseWriter, r *http.Request) {

}

func (p *postHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var postBody model.PostDTO
	if err = json.Unmarshal(body, &postBody); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	err = p.service.CreatePost(userId, postBody)
	if err != nil {
	}
}

func (p *postHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {

}

func (p *postHandler) DeletePost(w http.ResponseWriter, r *http.Request) {

}
