package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) SetMyUsername(u string, new string) (int64, error) {
	var id int64
	if err := db.c.QueryRow("SELECT id FROM users WHERE username=?", u).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return id, fmt.Errorf("user does not exist: %w", err)
		}
	}
	_, err := db.c.Exec("UPDATE users SET username=? WHERE username=?", new, u)
	if err != nil {
		var user User
		if errors.Is(db.c.QueryRow("SELECT id, username FROM users WHERE username=?", u).Scan(&user.UserId, &user.Username), sql.ErrNoRows) {
			return id, fmt.Errorf("user does not exist: %w", err)
		}
	}
	_, err = db.c.Exec("UPDATE profiles SET profileName=? WHERE profileName=?", new, u)
	if err != nil {
		var profile Profile
		if errors.Is(db.c.QueryRow("SELECT id, username FROM profiles WHERE profileName=?", u).Scan(&profile.ProfileId, &profile.ProfileName), sql.ErrNoRows) {
			return id, fmt.Errorf("user does not exist: %w", err)
		}
	}
	return id, nil
}
