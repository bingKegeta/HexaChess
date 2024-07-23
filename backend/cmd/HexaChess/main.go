package main

import (
	"context"
	"fmt"
	"net/http"

	db "github.com/BingKegeta/Hexachess/backend/DB"
	"github.com/BingKegeta/Hexachess/backend/internal/api"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: api.Routes(),
	}
	fmt.Println("Server started.")

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8000",
		Handler: a.router,
	}
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

func main() {
	db.InitDB("../../DB/hexachess.db")
	defer db.CloseDB()

	app := New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
