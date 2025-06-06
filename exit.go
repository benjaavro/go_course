package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("deferred statement")
	fmt.Println("Starting my go routine")

	os.Exit(1)

	fmt.Println("Exiting my go routine") // shouldn't be executed
}
