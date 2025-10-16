package routes

import (
	"net/http"
	"v1/monorepo/handlers"
)

const (
	Post        = "/post"
	PostByTitle = Post + "{title}"
	PostById    = Post + "{PostId}"
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
			URI:         Post,
			Method:      http.MethodGet,
			Function:    h.GetPosts,
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
