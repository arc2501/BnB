package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// Middleware func accepts a HTTP handler and returns the same
func WriteToConsole(next http.Handler) http.Handler {

	// returns a handler in form of Handler func which is a anonymous function
	// taking responseWriter and Request and passing them to next ServeHTTP
	// meanwhile just writing a thing on console
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// adds CSRF to all POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	// To create CSRF token
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// Loads and Save the session at every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
