package routes

import (
	"net/http"
	m "v1/monorepo/util/middleware"

	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// Puts all routes inside router
func Config(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, routeLogin)

	for _, route := range routes {
		if route.RequireAuth {
			r.HandleFunc(route.URI, m.Logger(m.IsAuth(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, m.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
