package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	db "github.com/BingKegeta/Hexachess/backend/DB"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-chi/render"
	_ "github.com/mattn/go-sqlite3"
)

type DBUser struct {
	ID       int    `json:"uid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterPlayer(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatal(err.Error())
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	res, err := db.CreateUser(user.Username, user.Password, user.Email)
	if err != nil {
		log.Fatal(err.Error())
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

type loginReq struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

func LoginPlayer(w http.ResponseWriter, r *http.Request) {
	var req loginReq
	var ret DBUser

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Fatal(err.Error())
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	result := db.GetUserByEmail(req.Identifier)
	err := result.Scan(&ret.ID, &ret.Username, &ret.Password, &ret.Email)

	if err != nil {
		log.Fatal(err.Error())
		render.Status(r, http.StatusNoContent)
		render.JSON(w, r, err.Error())
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(ret.Password), []byte(req.Password)); err != nil {
		// log.Fatal("Passwords do not Match!")
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, "Wrong Details!")
		return
	}

	// do login stuff
	render.Status(r, http.StatusOK)
	render.JSON(w, r, struct {
		ID int `json:"user_id"`
	}{ID: ret.ID})
}

func LogoutPlayer(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "LogoutPlayer endpoint")
}
