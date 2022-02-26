package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arc2501/bnb/pkg/config"
	"github.com/arc2501/bnb/pkg/handlers"
	"github.com/arc2501/bnb/pkg/render"
)

const portNumber = ":8080"

func main() {
	// create a local app variable
	var app config.AppConfig
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

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println("Starting Server at localhost", portNumber)
	http.ListenAndServe(portNumber, nil)

}
