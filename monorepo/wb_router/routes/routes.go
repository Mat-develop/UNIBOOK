package routes

import (
	"net/http"
	"v1/handlers"
	m "v1/util/middleware"

	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// will hava to add login handler after
// Puts all routes inside router
func Config(r *mux.Router, userHandler handlers.UserHandler, postHandler handlers.PostHandler, communityHandler handlers.CommunityHandler) *mux.Router {
	routes := GetUserRoutes(userHandler)
	routes = append(routes, routeLogin)
	routes = append(routes, GetPostRoutes(postHandler)...)
	routes = append(routes, GetCommunitiesRoutes(communityHandler)...)

	for _, route := range routes {
		if route.RequireAuth {
			r.HandleFunc(route.URI, m.Logger(m.IsAuth(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, m.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
