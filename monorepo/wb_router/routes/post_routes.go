package routes

import (
	"net/http"
	"v1/handlers"
)

const (
	Post      = "/post"
	Community = Post + "/c"

	PostById      = Post + "/{id}"
	PostByUserId  = Post + "/{userId}"
	CommunityPost = Community + "/{communityId}"
	PostByName    = Post + "/{PostId}"
	PostByTitle   = Post + "/{title}"
)

// dps faço o restante
func GetPostRoutes(h handlers.PostHandler) []Route {
	return []Route{
		{
			URI:         Post,
			Method:      http.MethodPost,
			Function:    h.CreatePost,
			RequireAuth: false,
		},
		{
			URI:         CommunityPost,
			Method:      http.MethodGet,
			Function:    h.GetCommunityPosts,
			RequireAuth: true,
		},
		{
			URI:         PostByUserId,
			Method:      http.MethodGet,
			Function:    h.GetUserPosts,
			RequireAuth: true,
		},
		{
			URI:         PostByTitle,
			Method:      http.MethodGet,
			Function:    h.GetPostByTittle,
			RequireAuth: true,
		},
		{
			URI:         PostById,
			Method:      http.MethodPut,
			Function:    h.UpdatePost,
			RequireAuth: true,
		},
		{
			URI:         PostById,
			Method:      http.MethodDelete,
			Function:    h.DeletePost,
			RequireAuth: true,
		},
	}

}
