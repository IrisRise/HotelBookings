package Render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	
	cnfg "github.com/IrisRise/HotelBookings/pkg/Config"
	"github.com/IrisRise/HotelBookings/pkg/Models"
)

var functions = template.FuncMap {

}

var app *cnfg.AppConfig

func NewTemplate(a *cnfg.AppConfig) {
	app = a
}

func AddDefaultData(data *Models.TemplateData)  *Models.TemplateData {

	return data
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *Models.TemplateData) error {

	var tc map[string]*template.Template
	var err error

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()

		if err != nil {
			log.Fatal(err)
		}
	}


	thisPage, ok := tc[tmpl]

	if !ok {
		
		log.Fatal("Could not get Template")
	}

	buff := new(bytes.Buffer)

	data = AddDefaultData(data)

	_ = thisPage.Execute(buff, data)

	_, err = buff.WriteTo(w)
	
	if err != nil {
		log.Fatal(err)

	}



	// parsedFiles, err := template.ParseFiles("./../../Templates/" + tmpl)

	// err = parsedFiles.Execute(w, nil)

	// if err != nil {
	// 	return err
	// }

	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./../../Templates/*.page.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)
		fmt.Println("Currently the page is", page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return nil, err
		}

		matches, err := filepath.Glob("./../../Templates/*.layout.html")

		if err != nil {
			return nil, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./../../Templates/*.layout.html")

			if err != nil {
				return nil, err
			}
		}

		templateCache[name] = ts
	}

	return templateCache, nil
}