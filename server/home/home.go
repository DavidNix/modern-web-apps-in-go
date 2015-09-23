package home

import "net/http"

func Hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gophers!\n"))
	})
}

func Goodbye() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("See ya later, Gophers!\n"))
	})
}
