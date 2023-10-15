package main

import (
	"log"
	"os"

	"github.com/CloudyKit/jet/v6"
)

type application struct {
	appName string
	srv     server
	debug   bool
	infoLog *log.Logger
	errLog  *log.Logger
	view    *jet.Set
}

type server struct {
	host string
	port int
	url  string
}

func main() {
	app := &application{
		appName: "HackerNews",
		srv: server{
			host: "localhost",
			port: 8090,
			url:  "http://localhost:8090",
		},
		debug:   true,
		infoLog: log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog:  log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Llongfile),
	}

	if app.debug {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"), jet.InDevelopmentMode())
	} else {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"))

	}

	if err := app.appServer(); err != nil {
		log.Fatal(err)
	}
}
