package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type CustomServer struct {
	server *http.Server
}

func NewServer(mux *chi.Mux) *CustomServer {
	return &CustomServer{
		server: &http.Server{
			Addr:           ":7777",
			Handler:        mux,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20, // 1 MB
		},
	}
}

func (customServer *CustomServer) Start() {
	log.Fatal(customServer.server.ListenAndServe())
}
