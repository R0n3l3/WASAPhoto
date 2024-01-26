package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) CreateUser(u string) (int64, error) {
	res, err := db.c.Exec("INSERT INTO users(username) VALUES (?)", u)
	id, _ := res.LastInsertId()
	if err != nil {
		var user User
		if err := db.c.QueryRow("SELECT id, username FROM users WHERE username=?", u).Scan(&user.UserId, &user.Username); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return id, fmt.Errorf("user does not exist: %w", err)
			}
		}
	}
	res, err = db.c.Exec("INSERT INTO profiles(profileName, photoNumber) VALUES (?, 0)", u)
	if err != nil {
		var profile Profile
		if err := db.c.QueryRow("SELECT id, username FROM profiles WHERE profileName=?", u).Scan(&profile.ProfileId, &profile.ProfileName); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return id, fmt.Errorf("profile does not exist: %w", err)
			}
		}
	}
	return id, nil
}
