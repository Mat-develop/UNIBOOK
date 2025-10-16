package routes

import (
	"net/http"
	"v1/monorepo/handlers"
)

var routeLogin = Route{
	URI:         "/login",
	Method:      http.MethodPost,
	Function:    handlers.Login,
	RequireAuth: false,
}
