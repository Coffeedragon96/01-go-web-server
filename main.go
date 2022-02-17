package main

import (
	"fmt"
	"log"
	"net/http"
)

// Hello Handler
// Two arguments, one is response writer
// Second one is the incoming request
func helloHandler(w http.ResponseWriter, r *http.Request) {

	// Checking if URL path is anything other than /hello
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Checking if request is anything other than GET
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusBadRequest)
		return
	}

	// Returning a simple response
	fmt.Fprintf(w, "hello!")
}

// Form Handler. Same arguments as above function
func formHandler(w http.ResponseWriter, r *http.Request) {

	// Check for correct path
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Parse the incoming form
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}

	// Notify if successful
	fmt.Fprintf(w, "POST request successful\n")

	// Retrieve individual values
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Print incoming values to command line
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	// We are telling the golang server to look in the filserver,
	// for the directory /static and look for index.html.
	// By default, the fileserver will look for index.html
	fileServer := http.FileServer(http.Dir("./static"))

	// Declare routes for declared URL pattern
	http.Handle("/", fileServer)

	// Declare routes for declared functions
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// Notify in command line for server initiation
	fmt.Printf("Starting server at port 8080\n")

	// Main actual server
	// ListenAndServe listens to all TCP connection requests directed towards the specific port
	// It always returns a non-nil err
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
