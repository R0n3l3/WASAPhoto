package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func (db *appdbimpl) SetMyUsername(u string, new string) error {

	id, err := db.GetUserId(u)
	log.Println("ecco l'id: " + strconv.FormatUint(id, 10))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	result, err := db.c.Exec("UPDATE users SET username=? WHERE userId=?", new, id)
	if err != nil {
		fmt.Println("Error updating row:", err)
		// Handle the error
	} else {
		rowsAffected, _ := result.RowsAffected()
		fmt.Println("Rows affected:", rowsAffected)
	}

	result, err = db.c.Exec("UPDATE profiles SET profileName=? WHERE profileName=?", new, u)
	if err != nil {
		if errors.Is(db.c.QueryRow("SELECT profileId FROM profiles WHERE profileName=?", u).Scan(), sql.ErrNoRows) {
			log.Println(err.Error())
			return err
		}
		log.Println(err.Error())
		return err
	}
	return nil
}
