package examples

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/DavidNix/mware"
	gorilla "github.com/gorilla/handlers"
)

func Hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Govna!"))
	})
}

func JsonExample() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"hello":           "world",
			"build a web app": "in Go",
			"so much":         "fun",
		})
	})
}

func JsonExampleMware() http.Handler {
	return mware.Build(JsonExample(), setContentType)
}

func JsonExampleHandlerFunc() http.Handler {
	return mware.BuildFunc(jsonHandlerFunc, logRequest, setContentType)
}

func jsonHandlerFunc(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"example": "HandlerFunc",
	})
}

func RecoveryExample() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		panic("oopsie")
	}
	return mware.BuildFunc(fn, logRequest, recovery)
}

// middlewares
//
func setContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		h.ServeHTTP(w, r)
	})
}

func logRequest(h http.Handler) http.Handler {
	return gorilla.LoggingHandler(os.Stderr, h)
}

func recovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("PANIC:", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Sorry, the server blew up."))
			}
		}()

		h.ServeHTTP(w, r)
	})
}
