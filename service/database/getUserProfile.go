package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) GetUserProfile(u string) (Profile, error) {
	var profile Profile
	err := db.c.QueryRow("SELECT profileId, profileName, photoNumber FROM profiles WHERE profileName= ?", u).Scan(&profile.ProfileId, &profile.ProfileName, &profile.PhotoNumber)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err)
		}
		return profile, err
	}
	return profile, nil
}
