package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NewRouter router for service
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		if route.IsSecured {
			handler = CheckTokenInCookieHandler(handler, route.Name)
		}

		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

//SecureRouter router with credentials checking
func SecureRouter(router *mux.Router) *mux.Router {

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		handler = CheckTokenInCookieHandler(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
