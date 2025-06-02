package main

import "fmt"

type Element struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Structs, enums and interfaces

	// snake_case
	// Strings, numbers, variables, file names, etc.

	// UPPERCASE
	// Constants

	// mixedCase
	// JS, HTML Document, variables, external libraries

	// lowercase
	// package names

	const MAXRETRIES = 5

	var employeeID = 1001

	fmt.Println("Max retries: ", MAXRETRIES)
	fmt.Println("Employee ID: ", employeeID)
}
