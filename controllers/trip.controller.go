package controllers

import "database/sql"

func DeleteTrip(db *sql.DB, id int) error {
	sqlStatement := `
		DELETE FROM Trip
		WHERE id = ?;`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}
