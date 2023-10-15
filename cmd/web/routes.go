package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)

	if app.debug {
		mux.Use(middleware.Logger)
	}
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := app.render(w, r, "index", nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	return mux
}
