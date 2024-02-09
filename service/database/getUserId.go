package database

func (db *appdbimpl) GetUserId(u string) (uint64, error) {
	var id uint64
	if err := db.c.QueryRow("SELECT userId FROM users WHERE username= ?", u).Scan(&id); err != nil {
		print(err.Error())
		return id, err
	}
	return id, nil
}
