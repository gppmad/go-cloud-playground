package main

import (
	"flag"
	"fmt"
	"os"
)

// Implement a simple CLI that accepts Name, Surname, Birthdate and save it to a database.
func main() {
	name := flag.String("name", "", "Your name")
	surname := flag.String("surname", "", "Your surname")

	// Print usage information if no flags are provided
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}

	flag.Parse()
	fmt.Println(*name, *surname)

	//GetDBVersion()
}
