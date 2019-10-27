package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	log.Printf("Starting server")
	routes := mux.NewRouter()
	routes.PathPrefix("/youtube/feed/").HandlerFunc(getYoutubeFeed).Methods("GET")
	routes.HandleFunc("/youtube/mp3/{id}", getYoutubeMP3Handler).Methods("GET")
	server := &http.Server{
		Handler:      routes,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
