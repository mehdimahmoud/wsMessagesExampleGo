package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

/*
 * GetRouter function
 * @return pointer to a mux.Router value
 */
func GetRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
