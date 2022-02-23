package main

import (
	cnfg "github.com/IrisRise/HotelBookings/pkg/Config"
	sc "github.com/IrisRise/HotelBookings/pkg/Handlers"
	rndr "github.com/IrisRise/HotelBookings/pkg/Render"
	"log"
	"net/http"
	"time"
	"github.com/alexedwards/scs/v2"
)

var app cnfg.AppConfig
var session *scs.SessionManager

func main() {

	
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session


	TemplateCache, err := rndr.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = TemplateCache
	app.UseCache = false

	Repo := sc.NewRepo(&app)
	sc.NewHandlers(Repo)

	rndr.NewTemplate(&app)


	// http.HandleFunc("/", sc.Repo.Home);
	// http.ListenAndServe(":8080", nil);

	server := &http.Server {
		Addr: ":8080",
		Handler: ChiRoutes(&app),
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}