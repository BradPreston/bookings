package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bradpreston/bookings/pkg/config"
	"github.com/bradpreston/bookings/pkg/models"
)

var functions = template.FuncMap {

}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(appConfig *config.AppConfig) {
	app = appConfig
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// RenderTemplate renders templates using html engine
func RenderTemplate(responseWriter http.ResponseWriter, html string, td *models.TemplateData) {
	var templateCache map[string]*template.Template

	if app.UseCache {
		// get the template cache from app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	template, ok := templateCache[html]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buffer := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = template.Execute(buffer, td)

	_, err := buffer.WriteTo(responseWriter)
	if err != nil {
		fmt.Println("error writing template to browswer:", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	mycache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return mycache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return mycache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return mycache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return mycache, err
			}
		}

		mycache[name] = ts
	}

	return mycache, nil
}