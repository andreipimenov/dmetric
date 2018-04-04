package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(a *Application) *chi.Mux {
	r := chi.NewRouter()
	r.Use(JSONCtx)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.MethodNotAllowed(NotAllowedHandler())
	r.NotFound(NotFoundHandler())

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", PingHandler())
		r.Post("/devices/{id}/metrics", MetricsHandler(a))
	})
	return r
}
