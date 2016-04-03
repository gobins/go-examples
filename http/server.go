package main

import (
	"io"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func Server(w http.ResponseWriter, r *http.Request) {
	log.Info("Received Request")
	io.WriteString(w, "Hi How are you!")
}

func main() {
	http.HandleFunc("/", Server)
	log.Info("Starting HTTP Server")
	http.ListenAndServe(":8080", nil)
}
