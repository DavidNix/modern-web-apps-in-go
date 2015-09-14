package httpd

import (
	"log"
	"net/http"
	"os"
)

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	// http.Handle("/", router.New())
	log.Println("http: Listening on port", port)
	http.ListenAndServe(":"+port, nil)
}
