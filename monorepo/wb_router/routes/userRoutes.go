package routes

import (
	"net/http"
	"v1/monorepo/handlers"
)

const (
	UriALL          = "/users"
	UriByID         = "/users/{userId}"
	UriFollowByID   = "/users/{userId}/follow"
	UriUnfollowByID = "/users/{userId}/unfollow"
)

func GetUserRoutes(h handlers.UserHandler) []Route {
	return []Route{
		{
			URI:         UriALL,
			Method:      http.MethodPost,
			Function:    h.CreateUser,
			RequireAuth: false,
		},
		{
			URI:         UriALL,
			Method:      http.MethodGet,
			Function:    h.GetUser,
			RequireAuth: true,
		},
		{
			URI:         UriByID,
			Method:      http.MethodPut,
			Function:    h.UpdateUser,
			RequireAuth: true,
		},
		{
			URI:         UriByID,
			Method:      http.MethodDelete,
			Function:    h.DeleteUser,
			RequireAuth: true,
		},
		{
			URI:         UriFollowByID,
			Method:      http.MethodPost,
			Function:    h.Follow,
			RequireAuth: true,
		},
		{
			URI:         UriUnfollowByID,
			Method:      http.MethodPost,
			Function:    h.Follow,
			RequireAuth: true,
		},
	}
}
