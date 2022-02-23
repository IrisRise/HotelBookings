package main

import (
	"github.com/IrisRise/HotelBookings/pkg/Config"
	"github.com/IrisRise/HotelBookings/pkg/Handlers"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func PatRoutes(app *Config.AppConfig) http.Handler {

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(Handlers.Repo.Home))

	return mux
}

func ChiRoutes(app *Config.AppConfig) http.Handler {

		mux := chi.NewRouter()

		mux.Use(middleware.Recoverer)
		mux.Use(NoSurf)
		mux.Use(SessionLoad)
	
		mux.Get("/", Handlers.Repo.Home)
		mux.Get("/about", Handlers.Repo.About)
	
		return mux
	
}