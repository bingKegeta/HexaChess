package db

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password, email string) (int64, error) {
	query := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`

	db := GetDB()

	hashedPwBy, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return -1, err
	}

	hashedPw := string(hashedPwBy)

	res, err := db.Exec(query, username, hashedPw, email)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func GetUserByEmail(identifier string) *sql.Row {
	query := `SELECT * FROM users WHERE email = ? OR username = ?`

	db := GetDB()

	result := db.QueryRow(query, identifier, identifier)
	return result
}
