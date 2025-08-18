package routes

import (
	"net/http"
	"v1/monorepo/handlers"
)

// TODO entender mais sobre isso aqui
// Esse depois vou refatorar para ser handler?
const (
	UriALL  = "/users"
	UriByID = "/users/{userId}"
)

var Handlers = handlers.NewUserHandler()

var userRoutes = []Route{
	{
		URI:         UriALL,
		Method:      http.MethodPost,
		Function:    Handlers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         UriALL,
		Method:      http.MethodGet,
		Function:    Handlers.GetUser,
		RequireAuth: true,
	},
	{
		URI:         UriByID,
		Method:      http.MethodPut,
		Function:    Handlers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:         UriByID,
		Method:      http.MethodDelete,
		Function:    Handlers.DeleteUser,
		RequireAuth: false,
	},
}
