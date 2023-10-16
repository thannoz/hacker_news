package main

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
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
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		app.session.Put(r.Context(), "test", "Carlos Konzo")
		err := app.render(w, r, "index", nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	mux.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		vars := make(jet.VarMap)
		vars.Set("test", app.session.GetString(r.Context(), "test"))
		err := app.render(w, r, "index", vars)
		if err != nil {
			log.Fatalln(err)
		}
	})

	return mux
}
