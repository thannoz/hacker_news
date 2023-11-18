package main

import "net/http"

func (app *application) LoadSession(next http.Handler) http.Handler {
	return app.session.LoadAndSave(next)
}

func (a *application) authRequired(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := a.session.GetInt(r.Context(), sessionKeyUserId)
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		w.Header().Add("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	}
}
