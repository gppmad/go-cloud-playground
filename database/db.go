package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetDBVersion() (string, error) {

	urlExample := "postgres://golang:golangdb@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", urlExample)
	if err != nil {
		fmt.Println("Connection error with the DB " + err.Error())
		return "", err
	}
	defer db.Close()

	// Perform a select query
	rows, err := db.Query("SELECT version();")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return "", err
	}
	defer rows.Close()

	var dbversion string
	// Print Result
	for rows.Next() {

		err := rows.Scan(&dbversion)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return "", err
		}
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return "", err
	}

	return dbversion, nil

}
