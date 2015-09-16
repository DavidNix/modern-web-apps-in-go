package routes

import (
	"modern-web-apps-in-go/server/examples"
	"modern-web-apps-in-go/server/router"
	"net/http"
)

// example curl:  curl -i localhost:8888/examples/1

func Routes() router.Routes {
	return router.Routes{
		{
			"examples_1",
			"GET", "/examples/1", prefixBody(prefixBody(examples.Hello())),
		}, {
			"examples_2",
			"GET", "/examples/2", examples.JsonExample(),
		}, {
			"examples_3",
			"GET", "/examples/3", examples.JsonExampleMware(),
		}, {
			"examples_4",
			"GET", "/examples/4", examples.JsonExampleHandlerFunc(),
		}, {
			"examples_5",
			"GET", "/examples/5", examples.RecoveryExample(),
		},
	}
}

func prefixBody(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Um "))
		h.ServeHTTP(w, r)
	})
}
