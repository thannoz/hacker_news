package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	_ "github.com/lib/pq"
	"github.com/upper/db/v4/adapter/postgresql"
)

type application struct {
	appName string
	srv     server
	debug   bool
	infoLog *log.Logger
	errLog  *log.Logger
	view    *jet.Set
	session *scs.SessionManager
}

type server struct {
	host string
	port int
	url  string
}

func main() {

	db, err := openDB("")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	upper, err := postgresql.New(db)

	if err != nil {
		log.Fatal(err)
	}
	defer func(upper) {
		err := upper.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(upper)

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

	// Initialize Jet template
	initJet(app)

	// Initialize Session
	initSession(app, db)

	if err := app.appServer(); err != nil {
		log.Fatal(err)
	}
}

func initSession(app *application, db *sql.DB) {
	app.session = scs.New()
	app.session.Lifetime = 24 * time.Hour
	app.session.Cookie.Persist = true
	app.session.Cookie.Name = app.appName
	app.session.Cookie.Domain = app.srv.host
	app.session.Cookie.SameSite = http.SameSiteStrictMode
	app.session.Store = postgresstore.New(db)
}

func initJet(app *application) {
	if app.debug {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"), jet.InDevelopmentMode())
	} else {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"))

	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
