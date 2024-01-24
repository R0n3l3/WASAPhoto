package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) GetUserId(u string) (int64, error) {
	var user User
	err := db.c.QueryRow("SELECT id, username FROM users WHERE username= ?", &user.UserId, &user.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.UserId, fmt.Errorf("user does not exist: %w", err)
		}
	}
	return user.UserId, nil
}
