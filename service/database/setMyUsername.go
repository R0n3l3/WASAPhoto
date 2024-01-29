package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) SetMyUsername(u string, new string) (User, error) {
	var user User
	var profile Profile
	id, _ := db.GetUserId(u)
	if err := db.c.QueryRow("UPDATE users SET username=? WHERE userId=?", new, id).Scan(&user); err != nil {
		if errors.Is(db.c.QueryRow("SELECT userId, username FROM users WHERE username=?", u).Scan(&user.UserId, &user.Username), sql.ErrNoRows) {
			return user, fmt.Errorf("user does not exist: %w", err)
		}
	}
	if err := db.c.QueryRow("UPDATE profiles SET profileName=? WHERE profileName=?", new, u).Scan(&profile); err != nil {
		if errors.Is(db.c.QueryRow("SELECT profileId, profileName FROM profiles WHERE profileName=?", u).Scan(&profile.ProfileId, &profile.ProfileName), sql.ErrNoRows) {
			return user, fmt.Errorf("profile does not exist: %w", err)
		}
	}
	return user, nil
}
