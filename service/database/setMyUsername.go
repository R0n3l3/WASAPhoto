package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) SetMyUsername(u string, new string) (int64, error) {
	id, _ := db.GetUserId(u)
	_, err := db.c.Exec("UPDATE users SET username=? WHERE userId=?", new, id)
	if err != nil {
		var user User
		if errors.Is(db.c.QueryRow("SELECT userId, username FROM users WHERE username=?", u).Scan(&user.UserId, &user.Username), sql.ErrNoRows) {
			return id, fmt.Errorf("user does not exist: %w", err)
		}
	}
	_, err = db.c.Exec("UPDATE profiles SET profileName=? WHERE profileName=?", new, u)
	if err != nil {
		var profile Profile
		if errors.Is(db.c.QueryRow("SELECT profileId, profileName FROM profiles WHERE profileName=?", u).Scan(&profile.ProfileId, &profile.ProfileName), sql.ErrNoRows) {
			return id, fmt.Errorf("profile does not exist: %w", err)
		}
	}
	return id, nil
}
