package routes

import (
	"net/http"
	"v1/monorepo/handlers"
)

const (
	UriALL            = "/users"
	UriByID           = "/users/{userId}"
	UriFollowByID     = "/users/{userId}/follow"
	UriUnfollowByID   = "/users/{userId}/unfollow"
	UriFollowers      = "/users/{userId}/followers"
	UriUpdatePassword = "/users/{userId}/update-password"
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
		{
			URI:         UriFollowers,
			Method:      http.MethodGet,
			Function:    h.Followers,
			RequireAuth: true,
		},
		{
			URI:         UriUpdatePassword,
			Method:      http.MethodGet,
			Function:    h.UpdatePassword,
			RequireAuth: true,
		},
	}
}
