package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

var server = "localhost"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

// addValues adds two intefers and return the sum
func addValues(x, y int) int {
	return x + y
}

//main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on %s%s", server, portNumber))
	_ = http.ListenAndServe(server+portNumber, nil)
}