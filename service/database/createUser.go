package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) CreateUser(u string) (int64, error) {
	res, err := db.c.Exec("INSERT INTO users(username) VALUES (?)", u)
	if err != nil {
		panic(err) //TODO
	}
	id, err := res.LastInsertId()
	if err != nil {
		var user User
		if err := db.c.QueryRow("SELECT userId FROM users WHERE userId=?", u).Scan(&user.UserId); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return id, fmt.Errorf("user does not exist: %w", err)
			}
		}
	}
	_, err = db.c.Exec("INSERT INTO profiles(profileName, photoNumber) VALUES (?, 0)", u)
	if err != nil {
		panic(err) //TODO
	}
	idProf, err := res.LastInsertId()
	if err != nil {
		var profile Profile
		if err := db.c.QueryRow("SELECT profileId FROM profiles WHERE profileId=?", u).Scan(&profile.ProfileId); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return idProf, fmt.Errorf("profile does not exist: %w", err)
			}
		}
	}
	if idProf != id {
		return 0, fmt.Errorf("error in id generation")
	}
	return id, nil
}
