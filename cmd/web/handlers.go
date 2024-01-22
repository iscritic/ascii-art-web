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

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	banners := []string{"shadow", "standard", "thinkertoy"}
	data := PageData{Banners: banners}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func AsciiArt(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/ascii-art" {
	// 	http.NotFound(w, r)
	// 	return
	// }

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	text := r.FormValue("text")
	selectedBanner := r.FormValue("banner")
	result, err := GenerateAscii(text, selectedBanner)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := PageData{Result: result}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
