package database

import (
	"database/sql"
	"log"
)

func (db *appdbimpl) GetUserProfiles(u string) ([]Profile, error) {
	var profiles []Profile
	statement := "SELECT profileId, profileName FROM profiles WHERE profiles.profileName='" + u + "'"
	rows, err := db.c.Query(statement)
	if err != nil {
		log.Println(err.Error())
		return profiles, err
	}
	defer rows.Close()
	for rows.Next() {
		var profile Profile
		if err := rows.Scan(&profile.ProfileId, &profile.ProfileName); err != nil {
			log.Println(err.Error())
			return profiles, err
		}
		profiles = append(profiles, profile)
	}
	if err := rows.Err(); err != nil {
		log.Println(err.Error())
		return profiles, err
	}
	if len(profiles) == 0 {
		return profiles, sql.ErrNoRows
	}
	return profiles, nil
}
