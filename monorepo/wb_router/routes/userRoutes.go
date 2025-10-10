package routes

import (
	"net/http"
	"v1/monorepo/handlers"
)

const (
	User              = "/users"
	UriByID           = User + "/{userId}"
	UriFollowByID     = User + "/{userId}/follow"
	UriUnfollowByID   = User + "/{userId}/unfollow"
	UriFollowers      = User + "/{userId}/followers"
	UriUpdatePassword = User + "/{userId}/update-password"
)

func GetUserRoutes(h handlers.UserHandler) []Route {
	return []Route{
		{
			URI:         User,
			Method:      http.MethodPost,
			Function:    h.CreateUser,
			RequireAuth: false,
		},
		{
			URI:         User,
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
