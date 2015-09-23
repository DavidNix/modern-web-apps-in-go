package router

import "github.com/gorilla/mux"

type Routes []Route

func New(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		// each method below returns *mux.Route
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route)
	}

	return router
}
