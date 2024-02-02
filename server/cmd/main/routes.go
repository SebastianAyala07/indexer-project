package main

import (
	"server/zinsearch/web"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(sh *web.SearchHandler) *chi.Mux {
	mux := chi.NewMux()
	mux.Use(
		middleware.Logger,
		middleware.Recoverer,
	)
	mux.Post("/search", sh.SearchHandler)
	return mux
}
