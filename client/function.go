package client

import (
	"html/template"
	"net/http"
)

var pagesDir = "client/pages/"

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		pagesDir + "Home.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
