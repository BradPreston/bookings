package main

import (
	"testing"

	"github.com/bradpreston/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRotues(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
		case *chi.Mux:
			// do nothing
		default:
			t.Errorf("Type is not *chi.Mux. Type is %T", v)
	}
}