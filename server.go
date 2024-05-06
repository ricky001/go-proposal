package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Define the directory containing the HTML file
	dir := "./"

	// Define the file name
	file := "index.html"

	// Create a file server handler
	fs := http.FileServer(http.Dir(dir))

	// Handle requests to the root URL ("/") by serving the HTML file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, file)
	})

	// Handle requests to other URLs by serving static files
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define the port to listen on
	port := ":8080"

	// Start the web server
	fmt.Printf("Server listening on port %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
