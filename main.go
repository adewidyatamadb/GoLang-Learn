package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

var server = "localhost"

//main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on %s%s", server, portNumber))
	_ = http.ListenAndServe(server+portNumber, nil)
}
