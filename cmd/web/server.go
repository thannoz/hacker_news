package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) appServer() error {
	host := fmt.Sprintf("%s:%d", app.srv.host, app.srv.port)

	srv := http.Server{
		Handler:     app.routes(),
		Addr:        host,
		ReadTimeout: 300 * time.Second,
	}

	app.infoLog.Printf("Server is listening on %s\n", host)

	return srv.ListenAndServe()
}
