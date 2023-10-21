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
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/thannoz/hackernews/models"
	"github.com/upper/db/v4"
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
	Models  models.Models
}

type server struct {
	host string
	port int
	url  string
}

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("error loading .env file")
	}

	db2, err := openDB(os.Getenv("DSN_CONN"))
	if err != nil {
		log.Fatal(err)
	}
	defer db2.Close()

	upper, err := postgresql.New(db2)

	if err != nil {
		log.Fatal(err)
	}
	defer func(upper db.Session) {
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
		Models:  models.NewModel(upper),
	}

	// Initialize Jet template
	initJet(app)

	// Initialize Session
	initSession(app, db2)

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
