package Handlers

import (
	"github.com/IrisRise/HotelBookings/pkg/Config"
	"github.com/IrisRise/HotelBookings/pkg/Models"
	spr "github.com/IrisRise/HotelBookings/pkg/Render"
	"net/http"
)

//Repository
type Repository struct {
	App *Config.AppConfig	
}

var Repo *Repository

func NewRepo(a *Config.AppConfig) *Repository {
	return &Repository {
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func(rp *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	rp.App.Session.Put(r.Context(), "Remote IP", remoteIP)

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	spr.RenderTemplate(w, "home.page.html", &Models.TemplateData{
		StringMap: stringMap,
	})
}

func(rp *Repository) About(w http.ResponseWriter, r *http.Request) {

	remoteIP := rp.App.Session.GetString(r.Context(), "Remote IP")
	

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again from " + remoteIP

	
	stringMap["Remote_IP"] = remoteIP

	spr.RenderTemplate(w, "about.page.html", &Models.TemplateData{
		StringMap: stringMap,
	})
}