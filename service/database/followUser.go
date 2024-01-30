package database

func (db *appdbimpl) FollowUser(toFollow string, follower string) (Profile, error) {
	idToFollow, _ := db.GetUserId(toFollow)
	idFollower, _ := db.GetUserId(follower)

	_, err := db.c.Exec("INSERT INTO follow(follower, following) VALUES (?, ?)", idFollower, idToFollow)
	if err != nil {
		panic(err) //TODO
	}

	profile, _ := db.GetUserProfile(toFollow)
	return profile, nil
}
