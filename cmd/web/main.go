package main

import (
	"fmt"
	"net/http"

	"github.com/adewidyatamadb/GoLang-Learn/pkg/handlers"
)

const portNumber = ":8080"

var server = "localhost"

//main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on %s%s", server, portNumber))
	_ = http.ListenAndServe(server+portNumber, nil)
}
