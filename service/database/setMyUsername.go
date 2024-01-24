package database

func (db *appdbimpl) SetMyUsername(u string, new string) (int64, error) {
	res, err := db.c.Exec("UPDATE users(?) ")
}
