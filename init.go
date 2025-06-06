package main

import "fmt"

func init() {
	fmt.Println("initializing package")
}

func init() {
	fmt.Println("initializing package for the second time")
}

func init() {
	fmt.Println("initializing package for the third time")
}

func main() {
	fmt.Println("main() started")
}
