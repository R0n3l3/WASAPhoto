package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) GetUserProfile(u string) (Profile, error) {
	var profile Profile

	if err := db.c.QueryRow("SELECT * FROM profiles WHERE profileName= ?", u).Scan(&profile); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return profile, fmt.Errorf("profile does not exist: %w", err)
		}
	}
	return profile, nil
}
