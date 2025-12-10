package handlers

import (
	"bytes"
	"net/http"
	"strings"
	"text/template"
	"ascii-art-web/server/ascii"
)

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "Bad request: missing fields", http.StatusBadRequest)
		return
	}

	ascii, err := ascii.Run(text, banner)
	if err != nil {
		// Detect exact type of error from Run()
		msg := err.Error()

		if strings.HasPrefix(msg, "400") {
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		if strings.HasPrefix(msg, "404") {
			http.Error(w, msg, http.StatusNotFound)
			return
		}

		// Any other error is internal server error
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := struct{ Result string }{Result: ascii}

	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, data); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)

}