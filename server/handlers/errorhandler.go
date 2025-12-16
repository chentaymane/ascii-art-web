package handlers

import (
	"html/template"
	"net/http"
)

func RenderError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)

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

	tmpl.Execute(w, data)
}
