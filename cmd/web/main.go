package main

import (
	"log"
	"net/http"
)

const port = ":7000"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/ascii-art", AsciiArt)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Starting server on http://127.0.0.1" + port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalln(err)
	}
}
