package handlers

import (
	"bytes"
	"net/http"
	"strings"
	"html/template"
	"ascii-art-web/server/ascii"
)

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RenderError(w, http.StatusBadRequest, "Bad request")
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		RenderError(w, http.StatusBadRequest, "Missing fields")
		return
	}

	asciiResult, err := ascii.Run(text, banner)
	if err != nil {
		msg := err.Error()

		if strings.HasPrefix(msg, "400") {
			RenderError(w, http.StatusBadRequest, msg)
			return
		}

		if strings.HasPrefix(msg, "404") {
			RenderError(w, http.StatusNotFound, msg)
			return
		}

		RenderError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	data := struct {
		Result string
	}{
		Result: asciiResult,
	}

	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		RenderError(w, http.StatusNotFound, "Template not found")
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		RenderError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	buf.WriteTo(w)
}
