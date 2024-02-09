package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) SetMyUsername(u string, new string) (User, error) {
	var user User

	id, err := db.GetUserId(u)
	if err != nil {
		print(err.Error())
		return user, err
	}

	if err := db.c.QueryRow("UPDATE users SET username=? WHERE userId=?", new, id).Scan(&user); err != nil {
		print(err.Error())
		return user, err
	}

	if err := db.c.QueryRow("UPDATE profiles SET profileName=? WHERE profileName=?", new, u).Scan(); err != nil {
		if errors.Is(db.c.QueryRow("SELECT profileId FROM profiles WHERE profileName=?", u).Scan(), sql.ErrNoRows) {
			print(err.Error())
			return user, err
		}
		print(err.Error())
		return user, err
	}
	return user, nil
}
