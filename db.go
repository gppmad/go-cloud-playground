package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetDBVersion() {

	urlExample := "postgres://golang:golangdb@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", urlExample)
	if err != nil {
		fmt.Println("Connection error with the DB " + err.Error())
		return
	}
	defer db.Close()

	// Perform a select query
	rows, err := db.Query("SELECT version();")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	// Print Result
	for rows.Next() {
		var dbversion string

		err := rows.Scan(&dbversion)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		fmt.Printf("DB Version: %s\n", dbversion)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return
	}
}
