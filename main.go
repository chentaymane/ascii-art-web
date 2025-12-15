package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web/server/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiHandler)
	http.Handle("/static/",
    http.StripPrefix("/static/",
        http.FileServer(http.Dir("templates")),
    ),
)

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Server failed:", err)
	}
}
