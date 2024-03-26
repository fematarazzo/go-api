package database

import (
	"database/sql"

	"api/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.DBStringConnection)
	if error != nil {
		return nil, error
	}

	if error := db.Ping(); error != nil {
		db.Close()
		return nil, error
	}

	return db, nil
}
