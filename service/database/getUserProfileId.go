package database

import "log"

func (db *appdbimpl) GetUserProfileId(id uint64) (Profile, error) {
	var profile Profile

	if err := db.c.QueryRow("SELECT profileId, profileName, photoNumber FROM profiles WHERE profileId= ?", id).Scan(&profile.ProfileId, &profile.ProfileName, &profile.PhotoNumber); err != nil {
		log.Println(err.Error())
		return profile, err
	}
	return profile, nil
}
