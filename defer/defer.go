package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("e:/test.txt")
	defer closeFile(f)
	writeFile(f)
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func createFile(s string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(s)
	if err != nil {
		panic(err)
	}
	return f
}
