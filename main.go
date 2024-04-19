package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

func contact_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Contact Page !")
}

func about_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the About Page !")
}


func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contact", contact_page)
	http.HandleFunc("/about", about_page)	
	
	http.ListenAndServe(":8080", nil)
}


func main() {
	handleRequest()
}

