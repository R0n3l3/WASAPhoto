package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) IsAuthorized(token uint64) bool {
	var tk uint64
	if err := db.c.QueryRow("SELECT token FROM auth WHERE token=?", token).Scan(&tk); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return false
	}
	return true
}
