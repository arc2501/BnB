package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/arc2501/bnb/internal/config"
	"github.com/arc2501/bnb/internal/handlers"
	"github.com/arc2501/bnb/internal/render"
)

const portNumber = ":8080"

// create a global app variable
var app config.AppConfig

// declare session so that in future can be used by middleware and all
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	// initializing a session variable
	session = scs.New()
	// defining its parameter
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// intitialize this above session on the app.Session
	app.Session = session
	// get the template cache from app config
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	// app holding template cache
	app.TemplateCache = tc
	app.UseCache = false

	// create a new repo object
	repo := handlers.NewRepo(&app)
	// pass it back to newhandler
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println("Starting Server at localhost", portNumber)
	err = srv.ListenAndServe()
	log.Fatal(err)

}
