package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/arc2501/bnb/internal/config"
	"github.com/arc2501/bnb/internal/models"
	"github.com/justinas/nosurf"
)

// This is for those things which we cannot do inside Go
// But Go allow us to define our own functions here.
// Like Fomatting a Date in certain way
var functions = template.FuncMap{}

// Declaring an app variable which we will use inside render
var app *config.AppConfig

// NewTemplates function sets the config for the template package
// in a way this is our entry of app from main to this pkg
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	// this is just for the developer mode
	// if USe cache is true then use it
	// otherwise rebuild it
	// and in main we will set it to false
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// having all the ts
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	// writing that ts content to the buffer
	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)
	_ = t.Execute(buf, td)
	// Making that buffer object write to the Response Writer
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to Browser", err)
	}

}

// Create a Template (set) cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	// Create a map as a cache
	myCache := map[string]*template.Template{}
	// have all pages here
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}
	// for every page
	for _, page := range pages {
		// Extract the page name
		name := filepath.Base(page)
		// create a new template in the memory
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// Look for Layout match
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		// if there is a layout
		if len(matches) > 0 {
			// Merge it with the page
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		// write the template set to the cache with the name as key
		myCache[name] = ts
	}
	return myCache, nil
}
