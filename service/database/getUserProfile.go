package database

func (db *appdbimpl) GetUserProfile(u string) (Profile, error) {
	var profile Profile

	if err := db.c.QueryRow("SELECT * FROM profiles WHERE profileName= ?", u).Scan(&profile); err != nil {
		print(err.Error())
		return profile, err
	}
	return profile, nil
}
