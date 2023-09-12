package main

import (
	"net/http"
)

func main() {
    fs := http.FileServer(http.Dir("static")) // Serve files from the "static" directory
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Start the HTTP server on a specific port (e.g., 8080)
    http.ListenAndServe(":8080", nil)
}