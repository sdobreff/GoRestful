package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		logger := NewLogger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(logger.LogHandler())

	}

	return router
}
