package main

import "fmt"

func main() {
	// Create a test for this.
	dbVersion, err := GetDBVersion()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(dbVersion)

	// Create a simple schema and perform a insert
}
