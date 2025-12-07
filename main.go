package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}


func Run(input string, banner string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("400: empty input")
	}

	// Allowed banners
	allowed := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}
	if !allowed[banner] {
		return "", fmt.Errorf("404: banner not found")
	}

	fontPath := "banners/" + banner + ".txt"

	// Normalize line breaks
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	content, err := os.ReadFile(fontPath)
	if err != nil {
		return "", fmt.Errorf("404: banner file missing")
	}

	fontTxt := strings.ReplaceAll(string(content), "\r\n", "\n")
	fontLines := strings.Split(fontTxt, "\n")

	final := ""

	for _, line := range lines {
		if line == "" {
			final += "\n"
			continue
		}

		runes := []rune(line)
		chars := make([][]string, len(runes))

		for i, char := range runes {
			if char < ' ' || char > '~' {
				return "", fmt.Errorf("400: invalid character")
			}

			// Compute ASCII char index in banner file
			index := int((char - ' ') * 9 + 1)
			if index+8 > len(fontLines) {
				return "", fmt.Errorf("500: corrupted banner file")
			}

			chars[i] = fontLines[index : index+8]
		}

		// Build ASCII-art lines
		for h := 0; h < 8; h++ {
			for _, c := range chars {
				final += c[h]
			}
			final += "\n"
		}
	}

	return final, nil
}


func asciiHandler(w http.ResponseWriter, r *http.Request) {

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

	ascii, err := Run(text, banner)
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

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
