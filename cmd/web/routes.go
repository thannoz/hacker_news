package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)
	mux.Use(app.LoadSession)

	if app.debug {
		mux.Use(middleware.Logger)
	}

	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/public", http.StripPrefix("/public", fileServer))

	return mux
}
