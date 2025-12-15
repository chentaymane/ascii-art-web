package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderError(w, http.StatusNotFound, "Page not found")
		return
	}
	if r.Method != http.MethodGet {
		RenderError(w, http.StatusBadRequest, "Bad request")
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderError(w, http.StatusNotFound, "Template not found")
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		RenderError(w, http.StatusInternalServerError, "Internal server error")
	}
}
