package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) GetFollowing(me string) ([]Profile, error) {
	var followers []Profile

	res, err := db.c.Query("SELECT p.profileId, p.profileName, p.photoNumber FROM profiles p, profiles p1, follow f WHERE p.profileId=f.following AND f.follower=p1.profileId AND p1.profileName=?", me)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return followers, err
	}
	defer func(rows *sql.Rows) {
		if closeErr := res.Close(); closeErr != nil {
			log.Println(closeErr.Error())
		}
	}(res)
	for res.Next() {
		var profile Profile
		err = res.Scan(&profile.ProfileId, &profile.ProfileName, &profile.PhotoNumber)
		if err != nil {
			log.Println(err.Error())
			return followers, err
		}
		followers = append(followers, profile)
	}
	if err = res.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return followers, nil
}
