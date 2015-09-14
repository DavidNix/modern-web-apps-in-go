package server

import (
	"log"
	"modern-web-apps-in-go/server/router"
	"net/http"
	"os"
)

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	http.Handle("/", router.New(Routes()))
	log.Println("http: Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
