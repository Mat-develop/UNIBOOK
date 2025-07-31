package middleware

import (
	"log"
	"net/http"
	"v1/monorepo/util/authentication"
	"v1/monorepo/util/response"
)

// writes request info
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}

}

// verifies if user requesting is authenticated
func IsAuth(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			response.Erro(w, http.StatusUnauthorized, err)
			return
		}
		nextFunc(w, r)
	}

}
