package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/BingKegeta/Hexachess/backend/internal/handlers"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(render.SetContentType(
		render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/api", func(r chi.Router) {
		r.Mount("/players", PlayerRoutes())
		r.Mount("/game", GameRoutes())
	})

	return router
}

func PlayerRoutes() http.Handler {
	router := chi.NewRouter()

	router.Post("/", handlers.RegisterPlayer)
	router.Post("/login", handlers.LoginPlayer)
	router.Post("/logout", handlers.LogoutPlayer)

	return router
}

func GameRoutes() http.Handler {
	router := chi.NewRouter()

	return router
}
