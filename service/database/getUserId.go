package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) GetUserId(u string) (int64, error) {
	var id int64
	if err := db.c.QueryRow("SELECT userId FROM users WHERE username= ?", u).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return id, fmt.Errorf("user does not exist: %w", err)
		}
	}
	return id, nil
}
