package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/goji/httpauth"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)
	for _, route := range routes {
		var handler http.Handler

		if (route.Secure){
			handler = httpauth.SimpleBasicAuth("admin", *adminPW)(route.HandlerFunc)
		} else {
			handler = route.HandlerFunc
		}

		handler = HTTPProvider(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
