package api

import (
	"net/http"

	"github.com/BingKegeta/Hexachess/backend/internal/game"
	"github.com/go-chi/render"
)

func GetBoard(w http.ResponseWriter, r *http.Request) {
	g := game.NewGame()

	if g.Board == nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, string("Not Created Right or some other error"))
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, g.Board)
}
