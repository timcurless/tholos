package service

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		r.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)

	}

	return r
}
