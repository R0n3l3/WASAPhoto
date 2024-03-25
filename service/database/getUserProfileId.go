package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) GetUserProfileId(id uint64) (Profile, error) {
	var profile Profile

	err := db.c.QueryRow("SELECT profileId, profileName, photoNumber FROM profiles WHERE profileId= ?", id).Scan(&profile.ProfileId, &profile.ProfileName, &profile.PhotoNumber)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return profile, err
	}
	return profile, nil
}
