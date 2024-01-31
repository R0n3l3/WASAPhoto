package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) GetUserProfile(u string) (Profile, error) {
	var profile Profile

	if err := db.c.QueryRow("SELECT * FROM profiles WHERE profileName= ?", u).Scan(&profile); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return profile, err
		}
		log.Fatal(err)
	}
	return profile, nil
}
