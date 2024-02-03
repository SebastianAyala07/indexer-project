package main

import (
	"server/zinsearch/web"

	"github.com/rs/cors"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(sh *web.SearchHandler) *chi.Mux {
	mux := chi.NewMux()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           10,
	})
	mux.Use(
		middleware.Logger,
		middleware.Recoverer,
		cors.Handler,
	)
	mux.Post("/search", sh.SearchHandler)
	return mux
}
