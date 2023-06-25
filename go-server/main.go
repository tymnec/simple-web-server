package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	// Just print to the screen
	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse Form() err: %v", err)
		return
	}

	fmt.Fprintf(w, "post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	// Root
	http.Handle("/", fileServer)

	// This function will handle the form
	http.HandleFunc("/form", formHandler)

	// This function will handle hello (Another Feature)
	http.HandleFunc("/hello", helloHandler)

	fmt.Print("Starting the server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
