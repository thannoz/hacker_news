package main

import (
	"fmt"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

// Global variables for our application
type TemplateData struct {
	URL             string
	IsAuthenticated bool
	AuthUser        string
	Flash           string
	Error           string
	CSRFToken       string
}

// defaultData sets global variables for our application
func (app *application) defaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.URL = app.srv.url
	return td
}

// render renders the templates
func (app *application) render(w http.ResponseWriter, r *http.Request, view string, vars jet.VarMap) error {
	td := &TemplateData{}

	td = app.defaultData(td, r)

	tp, err := app.view.GetTemplate(fmt.Sprintf("%s.html", view))
	if err != nil {
		return err
	}

	if err = tp.Execute(w, vars, td); err != nil {
		return err
	}
	return nil
}
