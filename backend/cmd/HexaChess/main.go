package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/BingKegeta/Hexachess/backend/internal/api"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.DefaultLogger)

	r.Get("/board", api.GetBoard)

	http.ListenAndServe(":8000", r)
}
