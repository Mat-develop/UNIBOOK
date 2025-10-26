package handlers

import (
	"net/http"
	"v1/monorepo/community/service"
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

func (c *communityHandler) CreateCommunity(w http.ResponseWriter, r *http.Request)       {}
func (c *communityHandler) GetCommunityByID(w http.ResponseWriter, r *http.Request)      {}
func (c *communityHandler) GetCommunityByName(w http.ResponseWriter, r *http.Request)    {}
func (c *communityHandler) ListCommunities(w http.ResponseWriter, r *http.Request)       {}
func (c *communityHandler) DeleteCommunity(w http.ResponseWriter, r *http.Request)       {}
func (c *communityHandler) FollowCommunity(w http.ResponseWriter, r *http.Request)       {}
func (c *communityHandler) GetCommunityFollowers(w http.ResponseWriter, r *http.Request) {}
