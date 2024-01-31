package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) IsAuthorized(token uint64) bool {
	if err := db.c.QueryRow("SELECT * FROM auth WHERE token=?", token).Scan(); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Fatal(err)
		}
		return false
	}
	return true
}
