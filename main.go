// package main

// import (
// 	"database/sql"
// 	"flag"
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/gppmad/gocloudplayground/database"
// )

// // Implement a simple CLI that accepts Name, Surname, Birthdate and save it to a database.
// func main() {
// 	name := flag.String("name", "", "Your name")
// 	surname := flag.String("surname", "", "Your surname")
// 	birthdate := flag.String("birthdate", "", "Birthdate in YYYY-MM-DD format (e.g., 1992-02-29)")

// 	// Print usage information if no flags are provided
// 	if len(os.Args) == 1 {
// 		flag.Usage()
// 		os.Exit(1)
// 	}

// 	flag.Parse()
// 	fmt.Println(*name, *surname)

// 	urlExample := "postgres://golang:golangdb@localhost:5432/golangdatabase?sslmode=disable"
// 	db, err := sql.Open("postgres", urlExample)
// 	if err != nil {
// 		fmt.Println("Error opening database:", err)
// 		return
// 	}
// 	defer db.Close()

// 	// Parse birthdate string into time.Time format
// 	var birthdateParsed time.Time
// 	if *birthdate != "" {
// 		birthdateParsed, err = time.Parse("2006-01-02", *birthdate)
// 		if err != nil {
// 			fmt.Println("Error parsing birthdate:", err)
// 			return
// 		}
// 	}

// 	err, res := database.AddPerson(db, *name, *surname, birthdateParsed)
// 	if res {
// 		fmt.Println("User has been successfully added")
// 	} else if err != nil {
// 		fmt.Println("Error", err)
// 	}

// }
