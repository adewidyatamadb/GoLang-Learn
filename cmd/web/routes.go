package main

import (
	"net/http"

	"github.com/adewidyatamadb/GoLang-Learn/pkg/config"
	"github.com/adewidyatamadb/GoLang-Learn/pkg/handlers"
	"github.com/bmizerany/pat"
)

func route(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
