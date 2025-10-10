package routes

import (
	"net/http"
	"v1/monorepo/handlers"
)

const (
	Post = "/post"
)

// dps faço o restante
func GetPostRoutes(h handlers.UserHandler) []Route {
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
	}

}
