package controller

import (
	"log"
	"net/http"
)

// Register initialize mux
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/", crud())
	return mux
}

// Web starts the web
func Web() {
	http.Handle("/ping", ping())
	http.Handle("/", crud())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
