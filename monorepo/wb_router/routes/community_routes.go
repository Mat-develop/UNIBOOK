package routes

import (
	"net/http"
	"v1/v1/handlers"
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
			URI:         CommunityPost,
			Method:      http.MethodGet,
			Function:    c.CreateCommunity,
			RequireAuth: true,
		},
	}

}
