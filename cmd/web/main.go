package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bradpreston/bookings/internal/config"
	"github.com/bradpreston/bookings/internal/handlers"
	"github.com/bradpreston/bookings/internal/models"
	"github.com/bradpreston/bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)

const port string = ":8080"
var app config.AppConfig
var session * scs.SessionManager

// main is the main application function
func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
	

	fmt.Println("Server running on", port)
	srv := &http.Server {
		Addr: port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	// What am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	return nil
}