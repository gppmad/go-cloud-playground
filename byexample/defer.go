// package main

// import (
// 	"fmt"
// 	"os"
// )

// func createFile(path string) *os.File {
// 	fmt.Println("Creating")
// 	f, err := os.Create(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return f
// }

// func writeFile(f *os.File, data string) {
// 	fmt.Println("Writing")
// 	fmt.Fprintln(f, data)
// }

// func closeFile(f *os.File) {
// 	fmt.Println("Closing")
// 	err := f.Close()
// 	if err != nil {
// 		fmt.Println("Cannot close the file")
// 		os.Exit(1)
// 	}
// }

// func main() {
// 	file := createFile("./tmp.txt")
// 	defer closeFile(file)
// 	writeFile(file, "Hello")
// }
