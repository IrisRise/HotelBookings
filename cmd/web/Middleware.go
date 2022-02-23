package main

import (
	"fmt"
	"net/http"
	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

//Prevents CSRF attack by creating a NoSurf token
//Adds CSRF protection on all POST requests
func NoSurf(next http.Handler) http.Handler {
	
	CSRFHandler := nosurf.New(next)
	
	CSRFHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return CSRFHandler
}

//Loads and saves Session on every request
func SessionLoad(next http.Handler) http.Handler {
	
	return session.LoadAndSave(next)

}