package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adewidyatamadb/GoLang-Learn/pkg/config"
	"github.com/adewidyatamadb/GoLang-Learn/pkg/handlers"
	"github.com/adewidyatamadb/GoLang-Learn/pkg/render"
)

const portNumber = ":8080"

var server = "localhost"

//main is the main application function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on %s%s", server, portNumber))
	_ = http.ListenAndServe(server+portNumber, nil)
}
