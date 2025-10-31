package routes

import (
	"net/http"
	"v1/v1/handlers"
)

const (
	Communities = "c/all"
)

func GetCommunitiesRoutes(c handlers.CommunityHandler) []Route {
	return []Route{
		{
			URI:         Post,
			Method:      http.MethodPost,
			Function:    c.GetCommunityByName,
			RequireAuth: false,
		},
		{
			URI:         Communities,
			Method:      http.MethodGet,
			Function:    c.ListCommunities,
			RequireAuth: true,
		},
	}

}
