package server

import (
	"github.com/didil/k8s-pod-info-api/server/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Get("/api/v1/info", handlers.Info)

	return r
}
