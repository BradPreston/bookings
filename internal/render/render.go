package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bradpreston/bookings/internal/config"
	"github.com/bradpreston/bookings/internal/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap {}

var app *config.AppConfig
var pathToTemplates = "./templates"

// NewTemplates sets the config for the template package
func NewTemplates(appConfig *config.AppConfig) {
	app = appConfig
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders templates using html engine
func RenderTemplate(w http.ResponseWriter, r *http.Request, html string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	template, ok := tc[html]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buffer := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = template.Execute(buffer, td)

	_, err := buffer.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browswer:", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	mycache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates))
	if err != nil {
		return mycache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return mycache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
		if err != nil {
			return mycache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
			if err != nil {
				return mycache, err
			}
		}

		mycache[name] = ts
	}

	return mycache, nil
}