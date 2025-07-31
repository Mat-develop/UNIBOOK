package routes

import (
	"net/http"
	"v1/monorepo/handlers"
)

// Esse depois vou refatorar para ser
const (
	UriALL  = "/users"
	UriByID = "/users/{userId}"
)

var userRoutes = []Route{
	{
		URI:         UriALL,
		Method:      http.MethodPost,
		Function:    handlers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         UriALL,
		Method:      http.MethodGet,
		Function:    handlers.GetUser,
		RequireAuth: true,
	},
	{
		URI:         UriByID,
		Method:      http.MethodPut,
		Function:    handlers.UpdateUser,
		RequireAuth: false,
	},
	{
		URI:         UriByID,
		Method:      http.MethodDelete,
		Function:    handlers.DeleteUser,
		RequireAuth: false,
	},
}
