package handlers

import (
	"net/http"
	"strings"
	"v1/community/service"
	"v1/util/response"
)

type CommunityHandler interface {
	CreateCommunity(w http.ResponseWriter, r *http.Request)
	GetCommunityByID(w http.ResponseWriter, r *http.Request)
	GetCommunityByName(w http.ResponseWriter, r *http.Request)
	ListCommunities(w http.ResponseWriter, r *http.Request)
	DeleteCommunity(w http.ResponseWriter, r *http.Request)
	FollowCommunity(w http.ResponseWriter, r *http.Request)
	GetCommunityFollowers(w http.ResponseWriter, r *http.Request)
}

type communityHandler struct {
	service service.CommunityService
}

func NewCommunityHandler(service service.CommunityService) CommunityHandler {
	return &communityHandler{service: service}
}

func (c *communityHandler) CreateCommunity(w http.ResponseWriter, r *http.Request)  {}
func (c *communityHandler) GetCommunityByID(w http.ResponseWriter, r *http.Request) {}

func (c *communityHandler) GetCommunityByName(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(r.URL.Query().Get("user"))
	community, err := c.service.GetCommunityByName(name)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, community)
}

func (c *communityHandler) ListCommunities(w http.ResponseWriter, r *http.Request) {
	communities, err := c.service.ListCommunities()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, communities)
}

func (c *communityHandler) DeleteCommunity(w http.ResponseWriter, r *http.Request)       {}
func (c *communityHandler) FollowCommunity(w http.ResponseWriter, r *http.Request)       {}
func (c *communityHandler) GetCommunityFollowers(w http.ResponseWriter, r *http.Request) {}
