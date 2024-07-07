package handlers

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/x-vneer/go-server/internal/middleware"
)

func Handler(r *chi.Mux) {
	// middleware
	r.Use(chiMiddleware.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		// middleware for /account route

		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
