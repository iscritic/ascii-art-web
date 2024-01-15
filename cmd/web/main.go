package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/ascii-art", AsciiArt)

	log.Println("Starting server on http://localhost:8000...")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
