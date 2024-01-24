package database

func (db *appdbimpl) BanUser(u string, ban string) (int64, error) {
	userId, _ := db.c.Exec("SELECT userId FROM users WHERE users.username == ?", u)
	bannedId, _ := db.c.Exec("SELECT userId FROM users WHERE users.username == ?", ban)
	res, err := db.c.Exec("INSERT INTO banned(user, banned) VALUES (?)", userId, bannedId)
	if err != nil {
		//TODO
	}
	result, _ := res.LastInsertId()
	return result, _
}
