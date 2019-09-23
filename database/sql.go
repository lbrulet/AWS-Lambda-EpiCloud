package database

import (
	"database/sql"
	"fmt"
	"os"
)

func InitDB() (*sql.DB, error) {
	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_ENDPOINT"), os.Getenv("DATABASE_NAME"),
	)

	db, err := sql.Open("mysql", dnsStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
