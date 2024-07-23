package db

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password, email string) error {
	query := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`

	db := GetDB()

	hashedPwBy, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	hashedPw := string(hashedPwBy)

	_, err = db.Exec(query, username, hashedPw, email)

	return err
}

func GetUserByEmail(email, password string) (*sql.Row, error) {
	query := `SELECT * FROM users WHERE email = ?`

	db := GetDB()

	result := db.QueryRow(query, email)
	return result, nil
}
