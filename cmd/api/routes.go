package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// check if service is still alive
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/", app.GetNotesList)
	mux.Post("/add", app.PostNote)
	mux.Patch("/update", app.UpdateNote)
	mux.Delete("/delete", app.DeleteNote)

	mux.Post("/get-note-by-id", app.GetNoteByID)

	return mux
}
