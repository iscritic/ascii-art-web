package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Result  string
	Banners []string
}

var templates = template.Must(template.ParseFiles("templates/index.html"))

func Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		banners := []string{"shadow", "standard", "thinkertoy"}
		data := PageData{Banners: banners}
		templates.ExecuteTemplate(w, "index.html", data)
	case http.MethodPost:
		text := r.FormValue("text")
		selectedBanner := r.FormValue("banner")
		result, err := GenerateAscii(text, selectedBanner)
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		data := PageData{Result: result}
		templates.ExecuteTemplate(w, "index.html", data)
	default:
		w.Header().Set("Allow", "POST, GET")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func AsciiArt(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
