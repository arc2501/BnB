package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/arc2501/bnb/pkg/config"
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

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc := app.TemplateCache

	// having all the ts
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	// writing that ts content to the buffer
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
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
