package database

import "log"

func (db *appdbimpl) GetUserProfile(u string) (Profile, error) {
	var profile Profile

	if err := db.c.QueryRow("SELECT profileId, profileName, photoNumber FROM profiles WHERE profileName= ?", u).Scan(&profile.ProfileId, &profile.ProfileName, &profile.PhotoNumber); err != nil {
		log.Println(err.Error())
		return profile, err
	}
	return profile, nil
}
