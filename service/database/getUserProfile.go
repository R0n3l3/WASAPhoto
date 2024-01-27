package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) GetUserProfile(u string) (int64, error) {
	var id int64
	if err := db.c.QueryRow("SELECT profileId FROM profiles WHERE profileName= ?", u).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return id, fmt.Errorf("profile does not exist: %w", err)
		}
	}
	return id, nil
}
