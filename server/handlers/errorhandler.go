package handlers

import (
	"bytes"
	"html/template"
	"net/http"
)

func RenderError(w http.ResponseWriter, status int, message string) {
	
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	data := struct {
		Status  int
		Message string
		}{
		Status:  status,
		Message: message,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Everything is fine → write the page
	w.WriteHeader(status)
	buf.WriteTo(w)
}
